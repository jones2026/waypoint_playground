project = "my-project"

app "web" {
    build {
        use "pack" {}
    }

    deploy {
        use "docker" {}
    }
}
