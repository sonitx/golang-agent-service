#!/usr/bin/env bash

# Script to create a new agent based on the sample implementation.
# Usage: ./create_agent.sh <AgentName>
# Example: ./create_agent.sh Sale

set -euo pipefail

if [[ $# -ne 1 ]]; then
  echo "Usage: $0 <AgentName>"
  exit 1
fi

AGENT_NAME="$1"
# Convert the provided name to lower‑case for file naming
AGENT_FILE_NAME=$(echo "$AGENT_NAME" | tr '[:upper:]' '[:lower:]')
# Build a unique key for the agent (lower‑case and prefixed)
AGENT_KEY="agent_${AGENT_FILE_NAME}"

# ---------------------------------------------------------------------------
# 1. Ensure the key does not already exist in configs/agents.yml
# ---------------------------------------------------------------------------
if grep -q "key: \"$AGENT_KEY\"" "configs/agents.yml"; then
  echo "Error: Agent key \"$AGENT_KEY\" already exists in configs/agents.yml"
  exit 1
fi

# ---------------------------------------------------------------------------
# 2. Append the new agent definition to configs/agents.yml
# ---------------------------------------------------------------------------
cat >> "configs/agents.yml" <<EOF
- name: "Agent $AGENT_NAME"
  key: "$AGENT_KEY"
  enable: true
  description: "Agent $AGENT_NAME"
EOF

echo "Added agent entry to configs/agents.yml"

# ---------------------------------------------------------------------------
# 3. Clone the sample implementation and rename identifiers
# ---------------------------------------------------------------------------
# ---------------------------------------------------------------------------
# 3. Clone the sample_agent folder and rename identifiers
# ---------------------------------------------------------------------------
SRC_DIR="agents/implements/sample_agent"
DEST_DIR="agents/implements/${AGENT_FILE_NAME}"

# Create destination folder
mkdir -p "$DEST_DIR"
# Copy all files
cp "$SRC_DIR"/* "$DEST_DIR"/

# Rename package name inside all .go files
for f in "$DEST_DIR"/*.go; do
  sed -i '' "s/package sampleagent/package ${AGENT_FILE_NAME}/g" "$f"
  # Replace struct name (SampleAgent) with <AgentName>Agent
  sed -i '' "s/SampleAgent/${AGENT_NAME}Agent/g" "$f"
  # Replace constructor name (NewSampleAgent) with New<AgentName>Agent
  sed -i '' "s/NewSampleAgent/New${AGENT_NAME}Agent/g" "$f"
done

echo "Created implementation folder $DEST_DIR with updated files"

# ---------------------------------------------------------------------------
# 4. Add the mapping in services/agent_service.go (initAgent function) before END AI AUTO GENERATED comment
# ---------------------------------------------------------------------------
SERVICE_FILE="services/agent_service.go"
TMP_FILE="${SERVICE_FILE}.tmp"

awk -v key="$AGENT_KEY" -v name="${AGENT_NAME}Agent" '
  /func \(s \*AgentService\) initAgent\(\) map\[string\]agents.AgentInterface \{/ {print; next}
  /\/\/ END AI AUTO GENERATED/ {
    print "\tmapAgent[\""key"\"] = implements.New"name"()"
    print $0
    next
  }
  {print}
' "$SERVICE_FILE" > "$TMP_FILE" && mv "$TMP_FILE" "$SERVICE_FILE"

echo "Updated initAgent mapping in $SERVICE_FILE"


# Update the package comment if needed (optional)

echo "Created implementation folder $DEST_DIR"

# ---------------------------------------------------------------------------
# 4. Add the mapping in services/agent_service.go (initAgent function)
# ---------------------------------------------------------------------------
SERVICE_FILE="services/agent_service.go"
TMP_FILE="${SERVICE_FILE}.tmp"

awk -v key="$AGENT_KEY" -v name="${AGENT_NAME}Agent" '\
  /func \(s \*AgentService\) initAgent\(\) map\[string\]agents.AgentInterface \{/ {print; next} \
  { if ($0 ~ /^[[:space:]]*mapAgent\["[^"]+"\] =/) {print; next} } \
  $0 == "}" { \
    # Insert the new mapping just before the closing brace of the function
    print "\tmapAgent[\""key"\"] = implements.New"name"()"; \
    print $0; next } \
  {print}' "$SERVICE_FILE" > "$TMP_FILE" && mv "$TMP_FILE" "$SERVICE_FILE"

echo "Updated initAgent mapping in $SERVICE_FILE"

# ---------------------------------------------------------------------------
# 5. Make the script itself executable (optional when creating)
# ---------------------------------------------------------------------------
chmod +x "$0"

echo "Agent \"$AGENT_NAME\" created successfully."
