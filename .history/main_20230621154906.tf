provider "todoist" {
    api_key = "ca571c54f7ef468da8edea1f6eba159b2f61081b"
}

resource "todoist_task" "terraformCloudPoC" {
    content = "Build DataDog provider using its API"
    completed = false
}