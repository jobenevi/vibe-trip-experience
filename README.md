# Vibe Trip Experience

Vibe Trip Experience is a project developed in Go that aims to provide a vibrant and unique travel experience. This repository contains the source code of the project, including environment setup, tests, and continuous integration.

## Table of Contents

- [Installation](#installation)
- [Usage](#usage)
- [Tests](#tests)
- [Continuous Integration](#continuous-integration)
- [Contribution](#contribution)
- [License](#license)

## Installation

To install and set up the project locally, follow the steps below:

1. Clone the repository:

    ```sh
    git clone https://github.com/jobenevi/vibe-trip-experience.git
    cd vibe-trip-experience
    ```

2. Install the dependencies:

    ```sh
    go mod download
    ```

## Usage

To run the project locally, use the command below:
```sh
go run
```

## Tests
To run the tests, use the command below:

## Continuous Integration
This repository uses GitHub Actions for continuous integration. The workflows are configured to run tests and automatic versioning.

1. CI Workflow
The CI workflow is triggered on push and pull request to the develop branch. It runs the tests and checks the code quality.

2. Versioning Workflow
The versioning workflow is triggered on push to the main branch. It increments the project version and creates a new Git tag.

## Contribution
Contributions are welcome! Feel free to open issues and pull requests. To contribute, follow the steps below:

1. Fork the repository.
2. Create a branch for your feature (git checkout -b feature/new-feature).
3. Commit your changes (git commit -m 'Add new feature').
4. Push to the branch (git push origin feature/new-feature).
5. Open a Pull Request.

## License

### Explanation

1. **Title and Description**:
   - The title of the project is followed by a brief description of what the project does.

2. **Table of Contents**:
   - An index to facilitate navigation through the `README`.

3. **Installation**:
   - Instructions to clone the repository and install the dependencies.

4. **Usage**:
   - Instructions to run the project locally.

5. **Tests**:
   - Instructions to run the tests of the project.

6. **Continuous Integration**:
   - Description of the CI and versioning workflows configured in GitHub Actions.

7. **Contribution**:
   - Instructions to contribute to the project.

8. **License**:
   - Information about the project's license.

### Conclusion

With this , users and contributors will have a clear view of how to install, use, test, and contribute to the `vibe-trip-experience` project.