var reporter = require('cucumber-html-reporter');

var options = {
    theme: 'hierarchy',
    jsonFile: 'test/report/cucumber_report.json',
    output: 'test/report/cucumber_report.html',
    reportSuiteAsScenarios: true,
    launchReport: true,
    columnLayout: 1,
    metadata: {
        "App Version": "0.3.4",
        "Test Environment": "STAGING",
        "Browser": "Chrome  54.0.2840.98",
        "Platform": "Windows 10",
        "Parallel": "Scenarios",
        "Executed": "Remote"
    }
};

reporter.generate(options);
process.exit();
