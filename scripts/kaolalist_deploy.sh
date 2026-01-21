#!/bin/bash
# KaolaList All-in-One Deployment Script
# This script installs dependencies, builds KaolaList, and initializes settings.

set -e

echo "🐨 KaolaList Deployment Started..."

# 1. Environment Check
if ! command -v go &> /dev/null; then
    echo "Go is not installed. Please install Go 1.21+."
    exit 1
fi

# 2. Build Backend
echo "🔨 Building KaolaList binary..."
go build -o KaolaList main.go

# 3. Initialize Configuration
if [ ! -f "config.json" ]; then
    echo "📝 Creating default config.json..."
    cat > config.json <<EOF
{
  "temp_dir": "data/temp",
  "log": {
    "enable": true,
    "name": "data/log/kaolalist.log"
  },
  "scheme": {
    "http_port": 5244
  }
}
EOF
fi

# 4. Success Message
echo "✅ KaolaList is ready!"
echo "🚀 Run it with: ./KaolaList server"
