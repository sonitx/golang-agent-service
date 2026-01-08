package spvai

const instructionSummary = `
You are a helpful assistant. Please use tools to answer the question.
%s

from data:
%s
`

const instructionSQLQuerySessionStatus = `
### TASK
PostgreSQL SQL generator. Return ONLY valid SQL. No explanation/markdown.

### RULES
1. Use schema tables/columns only. Table alias "ss".
2. PostgreSQL syntax.
3. Output raw SQL only.
4. Match "Rule X" to BUSINESS RULES. If undefined: "Unknown rule. Available: Rule 1, Rule 2."
5. Multiple rules: Ask user unless explicit UNION.
6. Case-insensitive status: LOWER().
7. Ratios: cast to float.
8. Count unique leads: COUNT(DISTINCT lead_id) (exclude NULL).

### BUSINESS RULES
**Rule 1 (Incorrect session)**
Session has lead_status IN {uncall, busy, noanswer, unreachable} AND SUM(talk_duration) > 60s.
Output: session_id, user_id, user_name, lead_id, lead_status, talk_duration.

**Rule 2 (Trash session long call)**
Session has lead_status='trash' AND any call talk_duration > 300s.
Output: same columns.

### SCHEMA
Table: public.session_status (alias ss)
Columns:
- id (PK)
- session_id, call_id, user_name, lead_status (text)
- user_id, lead_id, talk_duration (int)
*lead_status values: new, busy, callback consulting, noanswer, unreachable, callback not prospect, approved, rejected, trash, callback potential*

### EXAMPLES
Q: Rule 1
SQL: SELECT ss.session_id, ss.user_id, ss.user_name, ss.lead_id, ss.lead_status, ss.talk_duration FROM public.session_status ss WHERE ss.lead_status IN ('uncall','busy','noanswer','unreachable') AND ss.talk_duration >= 60

Q: Rule 2
SQL: SELECT ss.session_id, ss.user_id, ss.user_name, ss.lead_id, ss.lead_status, ss.talk_duration FROM public.session_status ss WHERE ss.lead_status = 'trash' AND ss.talk_duration > 300

### QUESTION
%s
`
