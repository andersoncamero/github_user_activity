# GitHub User Activity CLI

This CLI tool allows you to fetch and display recent activities of a GitHub user, such as comments, pull requests, and descriptions related to their repositories.

## Features

- Query any GitHub user's recent activities by providing their username.
- Displays relevant activity details, including:
  - Activity type.
  - User who performed the activity.
  - Associated repository.
  - Pull request information and additional descriptions.

## Prerequisites

- **Go**: Ensure Go is installed on your system. [Download Go](https://go.dev/dl/).
- **GitHub Access**: The project uses the GitHub API to retrieve user information.

## Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/your_username/github_user_activity.git
   cd github_user_activity
   ```

2. Install the necessary dependencies:

   ```bash
   go mod tidy
   ```

3. Build the program:

   ```bash
   go build -o github_activity
   ```

## Usage

1. Run the program from the terminal:

   ```bash
   ./github_activity
   ```

2. Enter the GitHub username when prompted:

   ```text
   github-username: octocat
   ```

3. The program will display the user's recent activities:

   ```text
   Fetching github activity...
   Type: PushEvent, Actor: octocat, Repo: octocat/Hello-World, Payload: { PR: <nil>, Descrp: "Added new README content." }
   ```

## Architecture

The project follows a well-structured approach using the following packages:

- **`internal/model`**: Defines the data structures, such as `Activity`.
- **`internal/usecase`**: Implements the logic for interacting with the GitHub API.

## Code Example

The main program logic is in `main.go`, where:

1. The username is requested.
2. A channel (`chan`) is used to handle asynchronous responses.
3. The user's activities are displayed in the console.

```go
fmt.Print("github-username: ")
_, err := fmt.Scanln(&input)
if err != nil {
	fmt.Println("invalid username error: ", err)
	return
}
activity, err := usecase.GetActivity(input)
// ...
```

## Contributing

Contributions are welcome! If you find an issue or have an improvement, please open an issue or a pull request in the repository.

## License

This project is licensed under the [MIT License](LICENSE).

## Author

Developed by **Anderson Rafael Camero Bajaire**.

You can view the project online at [Task Tracker Project](https://roadmap.sh/projects/github-user-activity).
