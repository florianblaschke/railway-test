module github.com/forsuxess/railway-test/importers/railway-importer

go 1.21

require (
	github.com/forsuxess/railway-test/jobs v0.0.0
	github.com/sirupsen/logrus v1.9.3
)

require golang.org/x/sys v0.0.0-20220715151400-c0bba94af5f8 // indirect

replace github.com/forsuxess/railway-test/jobs => ../../jobs
