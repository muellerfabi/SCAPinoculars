package main

import "github.com/jritter/openscap-report-publisher/reportrenderer"

func main() {
	reportrenderer.RenderReport("resources/reports/report1/arf.xml", "resources/reports/report1/report.html")
}
