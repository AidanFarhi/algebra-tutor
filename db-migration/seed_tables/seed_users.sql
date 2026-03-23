-- Users
INSERT INTO app_user (email) VALUES
('alice@example.com'),
('bob@example.com');

-- Courses
INSERT INTO course (title) VALUES
('Intro to Algebra');

-- Units
INSERT INTO unit (course_id, title, order_index) VALUES
(1, 'Linear Equations', 1),
(1, 'Quadratic Equations', 2);

-- Unit Components
INSERT INTO unit_component (unit_id, type, title, order_index) VALUES
(1, 'lesson', 'Solving for x', 1),
(1, 'quiz', 'Linear Quiz 1', 2),
(2, 'lesson', 'Intro to Quadratics', 1);

-- Explanations (for lesson component id = 1)
INSERT INTO explanation (unit_component_id, content, order_index) VALUES
(1, 'To solve for x, isolate the variable on one side.', 1),
(1, 'Use inverse operations to simplify the equation.', 2);

-- Interactive Elements
INSERT INTO interactive_element (unit_component_id, type, input_config, behavior_config, order_index) VALUES
(1, 'math_input', '{"placeholder": "Enter equation"}', '{"mode": "algebra"}', 1);

-- Questions (quiz component id = 2)
INSERT INTO question (
  unit_component_id,
  interactive_element_id,
  prompt,
  answer_type,
  input_config,
  validation_config,
  grading_config,
  order_index
) VALUES
(
  2,
  NULL,
  'Solve: 2x + 3 = 7',
  'equation',
  '{"format": "linear"}',
  '{"solution": "x=2"}',
  '{"points": 5}',
  1
),
(
  2,
  NULL,
  'What is 5 + 7?',
  'number',
  '{"type": "integer"}',
  '{"answer": 12}',
  '{"points": 2}',
  2
);

-- Attempt (user 1 taking quiz component 2)
INSERT INTO attempt (
  user_id,
  unit_component_id,
  type,
  status,
  score,
  max_score
) VALUES
(1, 2, 'quiz', 'submitted', 6, 7);

-- Attempt Questions (snapshot of questions)
INSERT INTO attempt_question (
  attempt_id,
  question_id,
  order_index,
  prompt_snapshot,
  input_config_snapshot,
  validation_config_snapshot,
  grading_config_snapshot,
  max_score
) VALUES
(
  1,
  1,
  1,
  'Solve: 2x + 3 = 7',
  '{"format": "linear"}',
  '{"solution": "x=2"}',
  '{"points": 5}',
  5
),
(
  1,
  2,
  2,
  'What is 5 + 7?',
  '{"type": "integer"}',
  '{"answer": 12}',
  '{"points": 2}',
  2
);

-- Attempt Responses
INSERT INTO attempt_response (
  attempt_question_id,
  answer_raw,
  answer_parsed,
  is_correct,
  score,
  feedback
) VALUES
(
  1,
  'x=2',
  '{"x": 2}',
  TRUE,
  5,
  'Correct!'
),
(
  2,
  '12',
  '12',
  TRUE,
  1,
  'Correct but partial credit rule applied.'
);