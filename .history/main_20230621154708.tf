provider "todoist" {
    api_key = ""
}

resource "todoist_task" "terraformCloudPoC" {
    content = "Build DataDog provider using its API"
    due
}