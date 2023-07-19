<div align="center">
<h1 align="center">
<img src="https://raw.githubusercontent.com/PKief/vscode-material-icon-theme/ec559a9f6bfd399b82bb44393651661b08aaf7ba/icons/folder-markdown-open.svg" width="100" />
<br>DDDGo
</h1>
<h3>◦ </h3>
<h3>◦ Developed with the software and tools listed below.</h3>

<p align="center">
<img src="https://img.shields.io/badge/Docker-2496ED.svg?style&logo=Docker&logoColor=white" alt="Docker" />
<img src="https://img.shields.io/badge/GitHub%20Actions-2088FF.svg?style&logo=GitHub-Actions&logoColor=white" alt="GitHub%20Actions" />

<img src="https://img.shields.io/badge/Go-00ADD8.svg?style&logo=Go&logoColor=white" alt="Go" />
<img src="https://img.shields.io/badge/Wire-000000.svg?style&logo=Wire&logoColor=white" alt="Wire" />
<img src="https://img.shields.io/badge/Markdown-000000.svg?style&logo=Markdown&logoColor=white" alt="Markdown" />
</p>
<a href="https://codecov.io/gh/program-world-labs/DDDGo" > 
 <img src="https://codecov.io/gh/program-world-labs/DDDGo/branch/master/graph/badge.svg?token=9V22UZ4HPO"/> 
 </a>
<img src="https://img.shields.io/github/languages/top/program-world-labs/DDDGo?style&color=5D6D7E" alt="GitHub top language" />
<img src="https://img.shields.io/github/languages/code-size/program-world-labs/DDDGo?style&color=5D6D7E" alt="GitHub code size in bytes" />
<img src="https://img.shields.io/github/commit-activity/m/program-world-labs/DDDGo?style&color=5D6D7E" alt="GitHub commit activity" />
<img src="https://img.shields.io/github/license/program-world-labs/DDDGo?style&color=5D6D7E" alt="GitHub license" />

</div>

---

## 📒 Table of Contents

- [📒 Table of Contents](#-table-of-contents)
- [📍 Overview](#-overview)
- [⚙️ Features](#️-features)
- [� Getting Started](#-getting-started)
  - [📦 Installation](#-installation)
  - [🧪 Running Tests](#-running-tests)
- [🗺 Roadmap](#-roadmap)
- [🤝 Contributing](#-contributing)
- [📄 License](#-license)
- [👏 Acknowledgments](#-acknowledgments)

---

## 📍 Overview

This project is an API service designed and implemented using Domain-Driven Design (DDD) principles. DDD is an approach to software development that greatly emphasizes the importance of understanding and modeling the domain. It aims to ease the creation of complex applications by connecting the related pieces of the software into an evolving model.

The API service is structured around business domains, each represented by a bounded context. The bounded context isolates the domain's models, making them specific to the context and independent from others. This isolation allows the model to evolve independently, reducing the risk of changes in one context affecting others.

The service is built with a layered architecture, typical in DDD:

- **Domain Layer**: This is the core of the software, containing business logic and types, which doesn't depend on other layers.
- **Application Layer**: This layer drives the workflow of the application, directing the domain layer, and is kept thin with logic that does not belong in the domain.
- **Infrastructure Layer**: This layer provides generic technical capabilities that support higher layers (message sending, persistence, drawing UI components, and so on).
- **Adapter Layer (or Presentation Layer)**: This layer is responsible for presenting information to the user and interpreting the user's commands.

This DDD approach ensures that the focus is on the core domain and domain logic, basing complex designs on a model, and initiating a creative collaboration between technical and domain experts to iteratively refine a conceptual model that addresses particular domain problems.

The API service is designed to be scalable, maintainable, and organized around the business domain, providing a strong foundation for further development and adjustments as the understanding of the domain evolves.

---

## ⚙️ Features

- **HTTP Support**: The service can handle HTTP requests and responses.
- **Behavior-Driven Development (BDD)**: The service is developed with a focus on behavior, making it more aligned with business requirements and easier to understand.
- **Event Sourcing**: The state of the business objects is determined by a sequence of events, providing a great audit trail and history that can enable various business insights.
- **SQL Transactions**: The service supports SQL transactions to ensure data consistency and integrity.
- **Anti-Breakdown**
- **Anti-Penetration**
- **Anti-Avalanche**

---

## 🚀 Getting Started

### 📦 Installation

1. Clone the DDDGo repository:

```sh
git clone https://github.com/program-world-labs/DDDGo
```

2. Change to the project directory:

```sh
cd DDDGo
```

3. Install the dependencies:

```sh
go mod tidy
```

### 🧪 Running Tests

```sh
make test
```

---

## 🗺 Roadmap

> - [x] `ℹ️ Role Implementation`
> - [ ] `ℹ️ Event Sourcing Implementation`
> - [ ] `ℹ️ User Implementation`
> - [ ] `ℹ️ Group Implementation`
> - [ ] `ℹ️ Wallet Implementation`

---

## 🤝 Contributing

Contributions are always welcome! Please follow these steps:

1. Fork the project repository. This creates a copy of the project on your account that you can modify without affecting the original project.
2. Clone the forked repository to your local machine using a Git client like Git or GitHub Desktop.
3. Create a new branch with a descriptive name (e.g., `new-feature-branch` or `bugfix-issue-123`).

```sh
git checkout -b new-feature-branch
```

4. Make changes to the project's codebase.
5. Commit your changes to your local branch with a clear commit message that explains the changes you've made.

```sh
git commit -m 'Implemented new feature.'
```

6. Push your changes to your forked repository on GitHub using the following command

```sh
git push origin new-feature-branch
```

7. Create a new pull request to the original project repository. In the pull request, describe the changes you've made and why they're necessary.
   The project maintainers will review your changes and provide feedback or merge them into the main branch.

---

## 📄 License

This project is licensed under the `ℹ️ INSERT-LICENSE-TYPE` License. See the [LICENSE](https://docs.github.com/en/communities/setting-up-your-project-for-healthy-contributions/adding-a-license-to-a-repository) file for additional info.

---

## 👏 Acknowledgments

---
