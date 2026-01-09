# Railway Test - Importer Setup

This is a test setup for deploying Go importers to Railway with a shared jobs package in a monorepo structure.

## Structure

```
railway-test/
├── jobs/                      # Shared jobs package
│   ├── job.go
│   └── go.mod
├── importers/
│   └── railway-importer/      # Example importer
│       ├── main.go
│       ├── go.mod
│       └── Dockerfile
├── railway.toml               # Railway configuration
└── README.md
```

## How It Works

- **Monorepo approach**: Shared `jobs` package and all importers in one repository
- **Independent deployments**: Each importer can be deployed separately
- **Local module references**: Uses Go's `replace` directive to reference local `jobs` package
- **Automatic deployments**: Railway watches for changes and auto-deploys
- **Native cron scheduling**: Each importer runs on its own cron schedule

## Local Development

### Prerequisites

- Go 1.21+
- Docker (for testing builds)
- Railway CLI (optional, for deployments)

### Running Locally

```bash
cd importers/railway-importer
go mod download
go run main.go
```

Expected output:
```json
{"client":"railway_test_client","job_id":301,"level":"info","msg":"railway importer completed","run_id":1,"success":true,"time":"2024-01-01T12:00:00Z","title":"senior software engineer (m/f/d)"}
```

### Testing the Docker Build

From the `railway-test` directory:

```bash
# Build the Docker image
docker build -f importers/railway-importer/Dockerfile -t railway-importer-test .

# Run the container
docker run --rm railway-importer-test
```

## Railway Deployment

### Prerequisites

1. Railway account (https://railway.app)
2. Railway CLI installed (optional but recommended)

```bash
# Install Railway CLI
npm i -g @railway/cli

# Or with Homebrew
brew install railway
```

### Initial Setup

#### Option 1: Deploy via GitHub (Recommended)

1. Push this code to a GitHub repository
2. Go to Railway dashboard (https://railway.app/dashboard)
3. Click "New Project" → "Deploy from GitHub repo"
4. Select your repository
5. Railway will detect the `railway.toml` and set up the service automatically
6. The importer will start running on the defined cron schedule

#### Option 2: Deploy via CLI

```bash
# Login to Railway
railway login

# Initialize project (from railway-test directory)
railway init

# Link to project
railway link

# Deploy
railway up
```

### Configuration

The `railway.toml` file configures the deployment:

```toml
[[services]]
name = "railway-importer"
rootDirectory = "importers/railway-importer"
dockerfilePath = "importers/railway-importer/Dockerfile"
dockerContext = "."
cronSchedule = "*/5 * * * *"  # Every 5 minutes
watch = ["jobs/**", "importers/railway-importer/**"]
```

### Adjusting the Cron Schedule

Edit `railway.toml` to change when the importer runs:

```toml
cronSchedule = "*/5 * * * *"   # Every 5 minutes
cronSchedule = "*/15 * * * *"  # Every 15 minutes
cronSchedule = "0 * * * *"     # Every hour at minute 0
cronSchedule = "30 2 * * *"    # Daily at 2:30 AM
cronSchedule = "0 9 * * 1"     # Every Monday at 9:00 AM
```

Cron format: `MIN HOUR DAY MONTH WEEKDAY`

### Viewing Logs

#### Via Railway Dashboard
1. Go to your project
2. Click on the service
3. View the "Deployments" or "Logs" tab

#### Via CLI
```bash
railway logs
```

### Environment Variables

Add environment variables in `railway.toml`:

```toml
[[services]]
name = "railway-importer"
# ... other config ...

[services.env]
LOG_LEVEL = "info"
API_TIMEOUT = "30s"
```

Or add them via the Railway dashboard (Settings → Variables).

## Adding More Importers

### 1. Create New Importer Directory

```bash
mkdir -p importers/importer-b
```

### 2. Create the Files

Copy from `railway-importer` as a template:

```bash
cp importers/railway-importer/main.go importers/importer-b/main.go
cp importers/railway-importer/go.mod importers/importer-b/go.mod
cp importers/railway-importer/Dockerfile importers/importer-b/Dockerfile
```

### 3. Update the Code

Edit `importers/importer-b/main.go` with your importer logic.

Update `importers/importer-b/go.mod`:
```go
module github.com/forsuxess/railway-test/importers/importer-b

// ... rest stays the same
```

### 4. Add to railway.toml

```toml
[[services]]
name = "importer-b"
rootDirectory = "importers/importer-b"
dockerfilePath = "importers/importer-b/Dockerfile"
dockerContext = "."
cronSchedule = "*/15 * * * *"
watch = ["jobs/**", "importers/importer-b/**"]
```

### 5. Deploy

Push to GitHub or run `railway up` - Railway will automatically detect and deploy the new service.

## Updating the Shared Jobs Package

When you update `jobs/job.go`:

1. Make your changes to the file
2. Commit and push (if using GitHub integration)
3. Railway automatically redeploys **all importers** that use the jobs package
4. This is expected behavior - if shared code changes, all consumers should update

## Troubleshooting

### Build fails with "cannot find module"

Make sure your `go.mod` has the correct replace directive:
```go
replace github.com/forsuxess/railway-test/jobs => ../../jobs
```

### Importer not running on schedule

1. Check Railway dashboard → Service → Cron Jobs tab
2. Verify the cron schedule syntax in `railway.toml`
3. Check service logs for errors

### "Context deadline exceeded" during build

Railway has build time limits. If builds are slow:
1. Reduce dependencies where possible
2. Use multi-stage builds (already implemented)
3. Cache Go modules properly (already implemented)

### Changes not triggering deployment

Check the `watch` paths in `railway.toml`:
```toml
watch = ["jobs/**", "importers/railway-importer/**"]
```

Make sure changed files match these patterns.

## Cost Estimation

Railway pricing (as of 2024):
- $5/month base subscription
- Plus usage-based pricing:
  - vCPU: ~$0.000463/vCPU-minute
  - RAM: ~$0.000231/GB-minute

**Example for this importer:**
- Runs every 5 minutes (288 times/day)
- Each run takes ~10 seconds
- Total runtime: ~48 minutes/day = 1,440 minutes/month
- Estimated cost: ~$2-3/month per importer

## Best Practices

1. **Keep jobs package focused**: Only shared code that multiple importers need
2. **Independent importers**: Each importer should be self-contained
3. **Proper logging**: Use structured logging (logrus JSON formatter)
4. **Error handling**: Importers should handle errors gracefully and log them
5. **Idempotency**: Design importers to be safe to run multiple times
6. **Monitoring**: Check Railway logs regularly for errors

## Next Steps

1. Test the deployment with `railway-importer`
2. Add your actual import logic
3. Create additional importers as needed
4. Set up monitoring/alerting (Railway integrations)
5. Configure any required environment variables
6. Adjust cron schedules based on your needs

## Support

- Railway Documentation: https://docs.railway.app
- Railway Discord: https://discord.gg/railway
- Railway Status: https://status.railway.app