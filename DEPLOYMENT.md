# Railway Deployment Guide

Quick guide to deploy your importers to Railway.

## Prerequisites

1. ✅ Railway account: https://railway.app
2. ✅ GitHub account (for automatic deployments)
3. ✅ This code in a GitHub repository

## Deployment Steps

### Step 1: Push to GitHub

```bash
# If not already in git
cd railway-test
git init
git add .
git commit -m "Initial railway-test setup"

# Push to GitHub
git remote add origin https://github.com/yourusername/your-repo.git
git push -u origin main
```

### Step 2: Connect to Railway

1. Go to https://railway.app/dashboard
2. Click **"New Project"**
3. Select **"Deploy from GitHub repo"**
4. Authorize Railway to access your GitHub account
5. Select your repository
6. Railway will automatically detect `railway.toml` and configure services

### Step 3: Verify Deployment

1. Railway will show your service(s) in the dashboard
2. Click on **"railway-importer"**
3. Go to **"Deployments"** tab to see build progress
4. Once deployed, check **"Logs"** tab to see execution logs

### Step 4: Verify Cron Execution

The importer runs every 5 minutes by default. Wait a few minutes and check logs:

```json
{"client":"railway_test_client","job_id":301,"level":"info","msg":"railway importer completed","run_id":1,"success":true,"time":"...","title":"senior software engineer (m/f/d)"}
```

## Making Changes

### Update Importer Code

```bash
# Edit the importer
vim importers/railway-importer/main.go

# Commit and push
git add .
git commit -m "Update railway-importer logic"
git push
```

Railway automatically rebuilds and redeploys when you push.

### Update Shared Jobs Package

```bash
# Edit shared jobs
vim jobs/job.go

# Commit and push
git add .
git commit -m "Update jobs package"
git push
```

**Note:** All importers using the jobs package will be redeployed automatically.

### Change Cron Schedule

Edit `railway.toml`:

```toml
[[services]]
name = "railway-importer"
# ... other config ...
cronSchedule = "0 * * * *"  # Change to every hour
```

Commit and push - Railway will update the schedule.

## Add More Importers

### 1. Create New Importer

```bash
# Copy template
cp -r importers/railway-importer importers/importer-b

# Update the code
vim importers/importer-b/main.go

# Update go.mod
sed -i '' 's/railway-importer/importer-b/g' importers/importer-b/go.mod
```

### 2. Add to railway.toml

```toml
[[services]]
name = "importer-b"
rootDirectory = "importers/importer-b"
dockerfilePath = "importers/importer-b/Dockerfile"
dockerContext = "."
cronSchedule = "*/15 * * * *"  # Every 15 minutes
watch = ["jobs/**", "importers/importer-b/**"]
```

### 3. Deploy

```bash
git add .
git commit -m "Add importer-b"
git push
```

Railway automatically deploys the new service!

## Environment Variables

### Via railway.toml

```toml
[[services]]
name = "railway-importer"
# ... other config ...

[services.env]
API_KEY = "your-api-key"
LOG_LEVEL = "debug"
```

### Via Railway Dashboard (Recommended for Secrets)

1. Go to your service in Railway dashboard
2. Click **"Variables"** tab
3. Click **"New Variable"**
4. Add `API_KEY` = `your-secret-key`
5. Service will automatically redeploy

Access in Go:
```go
import "os"

apiKey := os.Getenv("API_KEY")
```

## Monitoring

### View Logs

**Via Dashboard:**
1. Click on service
2. Go to **"Logs"** tab
3. See real-time logs

**Via CLI:**
```bash
# Install Railway CLI
npm i -g @railway/cli

# Login
railway login

# Link to project
railway link

# View logs
railway logs
```

### Cron Execution Status

1. Go to service in Railway dashboard
2. Look for **"Cron"** section
3. Shows last execution time and status

## Troubleshooting

### Build Fails

**Check build logs:**
1. Go to service → Deployments
2. Click on failed deployment
3. Check build logs for errors

**Common issues:**
- Missing dependencies: Run `go mod tidy` locally first
- Wrong paths: Ensure `dockerContext = "."` points to railway-test root
- Module issues: Verify `replace` directive in go.mod

### Importer Not Running

**Check cron configuration:**
1. Verify `cronSchedule` syntax in railway.toml
2. Check Railway dashboard → Service → Cron section
3. Ensure service is deployed (not just built)

**Cron format:** `MIN HOUR DAY MONTH WEEKDAY`
- `*/5 * * * *` = Every 5 minutes
- `0 * * * *` = Every hour
- `30 2 * * *` = Daily at 2:30 AM

### Logs Show Errors

**Common errors:**

**"cannot find module"**
- Check go.mod replace directive
- Ensure jobs package is in correct location

**"connection refused"**
- Check if external APIs are accessible
- Verify environment variables are set

**"context deadline exceeded"**
- Increase timeout in your code
- Check if external services are responding

## Cost Management

### Check Usage

1. Go to Railway dashboard
2. Click **"Usage"** at the top
3. See current month's usage and costs

### Estimate Costs

For each importer:
- Runs every 5 minutes = 288 times/day
- Each run ~10 seconds = 48 minutes/day
- Monthly: ~1,440 minutes
- Cost: ~$2-3/month per importer

**Total for 10 importers: $5 base + $20-30 usage = $25-35/month**

### Optimize Costs

1. **Adjust schedules:** Run less frequently if possible
2. **Optimize code:** Faster execution = lower cost
3. **Combine jobs:** Group related imports if appropriate

## Rollback

### Rollback to Previous Deployment

1. Go to service → Deployments
2. Find previous successful deployment
3. Click **"⋯"** menu → **"Redeploy"**

### Rollback via Git

```bash
# Find the commit to rollback to
git log

# Revert to previous commit
git revert HEAD
git push
```

Railway automatically deploys the reverted version.

## Best Practices

1. ✅ **Test locally** before pushing
2. ✅ **Use environment variables** for secrets
3. ✅ **Monitor logs** regularly
4. ✅ **Set up alerts** (Railway integrations)
5. ✅ **Keep cron schedules reasonable** (avoid every minute)
6. ✅ **Handle errors gracefully** in your code
7. ✅ **Use structured logging** (JSON format)

## Quick Reference

### Railway CLI Commands

```bash
railway login              # Login to Railway
railway link               # Link to project
railway up                 # Deploy manually
railway logs               # View logs
railway variables          # Manage environment variables
railway status             # Check project status
```

### Useful Links

- Railway Dashboard: https://railway.app/dashboard
- Railway Docs: https://docs.railway.app
- Railway Status: https://status.railway.app
- Railway Discord: https://discord.gg/railway

## Next Steps

1. ✅ Deploy railway-importer
2. ✅ Verify cron execution in logs
3. ✅ Add your actual import logic
4. ✅ Create additional importers
5. ✅ Set up monitoring/alerts
6. ✅ Configure environment variables
7. ✅ Adjust cron schedules as needed