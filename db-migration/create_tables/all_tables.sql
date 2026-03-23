CREATE TABLE app_user (
  id SERIAL PRIMARY KEY,
  email TEXT UNIQUE NOT NULL,
  created_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE course (
  id SERIAL PRIMARY KEY,
  title TEXT NOT NULL,
  created_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE unit (
  id SERIAL PRIMARY KEY,
  course_id INT NOT NULL REFERENCES course(id) ON DELETE CASCADE,
  title TEXT NOT NULL,
  order_index INT NOT NULL,
  created_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE unit_component (
  id SERIAL PRIMARY KEY,
  unit_id INT NOT NULL REFERENCES unit(id) ON DELETE CASCADE,
  type TEXT NOT NULL CHECK (type IN ('lesson', 'quiz', 'test', 'exercise')),
  title TEXT,
  order_index INT NOT NULL,
  created_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE explanation (
  id SERIAL PRIMARY KEY,
  unit_component_id INT NOT NULL REFERENCES unit_component(id) ON DELETE CASCADE,
  content TEXT NOT NULL,
  order_index INT NOT NULL
);

CREATE TABLE interactive_element (
  id SERIAL PRIMARY KEY,
  unit_component_id INT NOT NULL REFERENCES unit_component(id) ON DELETE CASCADE,
  type TEXT NOT NULL,
  input_config JSONB NOT NULL DEFAULT '{}'::jsonb,
  behavior_config JSONB NOT NULL DEFAULT '{}'::jsonb,
  order_index INT NOT NULL
);

CREATE TABLE question (
  id SERIAL PRIMARY KEY,
  unit_component_id INT NOT NULL REFERENCES unit_component(id) ON DELETE CASCADE,
  interactive_element_id INT NULL REFERENCES interactive_element(id),  -- optional link to a widget
  prompt TEXT NOT NULL,
  answer_type TEXT NOT NULL CHECK (
    answer_type IN (
      'multiple_choice',
      'number',
      'expression',
      'equation',
      'set',
      'function',
      'graph'
    )
  ),
  input_config JSONB NOT NULL DEFAULT '{}'::jsonb,
  validation_config JSONB NOT NULL DEFAULT '{}'::jsonb,
  grading_config JSONB NOT NULL DEFAULT '{}'::jsonb,
  order_index INT NOT NULL,
  created_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE attempt (
  id SERIAL PRIMARY KEY,
  user_id INT NOT NULL REFERENCES app_user(id) ON DELETE CASCADE,
  unit_component_id INT NOT NULL REFERENCES unit_component(id) ON DELETE CASCADE,
  type TEXT NOT NULL CHECK (type IN ('quiz', 'test', 'exercise')),
  status TEXT NOT NULL CHECK (status IN ('in_progress', 'submitted', 'graded')),
  score NUMERIC,
  max_score NUMERIC,
  started_at TIMESTAMP DEFAULT NOW(),
  submitted_at TIMESTAMP
);

CREATE TABLE attempt_question (
  id SERIAL PRIMARY KEY,
  attempt_id INT NOT NULL REFERENCES attempt(id) ON DELETE CASCADE,
  question_id INT NOT NULL,
  order_index INT NOT NULL,
  prompt_snapshot TEXT NOT NULL,
  input_config_snapshot JSONB NOT NULL,
  validation_config_snapshot JSONB NOT NULL,
  grading_config_snapshot JSONB NOT NULL,
  max_score NUMERIC NOT NULL
);

CREATE TABLE attempt_response (
  id SERIAL PRIMARY KEY,
  attempt_question_id INT NOT NULL REFERENCES attempt_question(id) ON DELETE CASCADE,
  answer_raw TEXT,
  answer_parsed JSONB,
  is_correct BOOLEAN,
  score NUMERIC,
  feedback TEXT,
  created_at TIMESTAMP DEFAULT NOW()
);

CREATE INDEX idx_unit_course ON unit(course_id);
CREATE INDEX idx_component_unit ON unit_component(unit_id);
CREATE INDEX idx_question_component ON question(unit_component_id);

CREATE INDEX idx_attempt_user ON attempt(user_id);
CREATE INDEX idx_attempt_component ON attempt(unit_component_id);

CREATE INDEX idx_attempt_question_attempt ON attempt_question(attempt_id);
CREATE INDEX idx_response_attempt_question ON attempt_response(attempt_question_id);