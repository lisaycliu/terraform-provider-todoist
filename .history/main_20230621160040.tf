provider "todoist"{}

resource "todoist_task" "terraformCloudPoC" {
    content = "Build DataDog provider using its API"
    completed = false
}