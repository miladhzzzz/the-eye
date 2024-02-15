# Observability Platform

This project is an observability platform designed to provide deep insights into Kubernetes clusters using eBPF (extended Berkeley Packet Filter) technology. eBPF allows us to efficiently instrument and collect data at various levels of the Linux kernel, providing a powerful and flexible foundation for monitoring and observability.

## Features

- **Agent**: Collects metrics using eBPF and exposes them via an HTTP endpoint.
- **CLI Tool**: A command-line interface for deploying and configuring the observability platform.
- **Docker Compose**: Simplifies local development with Docker Compose.
- **Kubernetes Deployment**: Provides configurations for deploying the platform to a Kubernetes cluster.
- **Build Automation**: GitHub Actions workflows for automating the build and deployment processes.

## The Power of eBPF

[eBPF](https://ebpf.io/) is a revolutionary technology that allows us to extend the functionality of the Linux kernel in a safe and efficient manner. It's like having programmable hooks inside the kernel that enable us to inspect and manipulate various events and data. This observability platform leverages eBPF to gain unprecedented insights into the inner workings of a Kubernetes cluster.

### Story Behind eBPF Integration

Imagine trying to monitor a highly dynamic and complex system like Kubernetes. Traditional monitoring approaches often involve heavy instrumentation or reliance on external agents. We wanted something more lightweight, efficient, and capable of providing granular insights without adding significant overhead.

Enter eBPF.

The decision to use eBPF wasn't just about adopting a trendy technology; it was about addressing specific challenges in monitoring Kubernetes. We needed a solution that could:

1. **Low Overhead**: Kubernetes environments are highly dynamic, with containers spinning up and down rapidly. Traditional monitoring tools could introduce significant overhead. eBPF, being part of the kernel, operates at a lower level, minimizing the impact on system resources.

2. **Granular Visibility**: We wanted to go beyond basic metrics and gain deep insights into the interactions between containers, nodes, and the kernel itself. eBPF provides the ability to attach probes to specific kernel functions, allowing us to capture fine-grained data.

3. **Flexibility**: Kubernetes is an ever-evolving ecosystem, and we needed a monitoring solution that could adapt to changes. eBPF's programmability and flexibility enable us to customize our observability strategy without major overhauls.
For more detailed information, refer to individual directories in the repository.

## Technical Considerations

**Agent**
The agent component of this observability platform is responsible for running eBPF programs and exposing the collected metrics via an HTTP endpoint. Let's delve into some technical details:

**eBPF Programs**
In the ebpf/ directory, you'll find the monitor.go file, which contains the eBPF programs written in Go. These programs are attached to specific kernel events and functions, allowing us to capture data at different levels of the system.

**HTTP Endpoint**
The agent's main functionality is to provide an HTTP endpoint (/metrics) for exposing the collected metrics. This allows users to easily scrape and visualize the data using their preferred monitoring tools.

**CLI Tool**
The CLI tool, found in the cli/ directory, is the user interface for deploying and configuring the observability platform. It interacts with the Kubernetes API and orchestrates the deployment of the agent.

**Local Development**
For local development, we use Docker Compose to bring up the entire observability platform. The docker-compose.yml file defines the services, including the agent and any required dependencies.

**Kubernetes Deployment**
The k8s/ directory contains Kubernetes manifests for deploying the observability platform to a Kubernetes cluster. The agent-deployment.yml file specifies the deployment configuration for the agent, while the agent-service.yml file defines a service for exposing the agent internally.

**Build Automation**
To streamline the build and deployment processes, we've set up GitHub Actions workflows in the .github/workflows/ directory. These workflows, such as build-agent-image.yml and build-client-binary.yml, automate the building and availability of the agent image and client binary upon pushing changes to the main branch.

## Local

Prerequisites:

- Docker
- Docker Compose

To start the observability platform locally, use the following command:

```bash
make run-docker-compose
This command will bring up the agent and any necessary dependencies.
```

## Deployment

**Kubernetes Deployment**
Use the provided Kubernetes manifests in the k8s/ directory for deploying the platform to a Kubernetes cluster. Adjust the configurations as needed for your environment.

## GitHub Actions

The project includes GitHub Actions workflows (build-agent-image.yml and build-client-binary.yml) for automating the build and deployment processes. These workflows trigger on pushes to the main branch.

## Building

**Agent Image**
To build the agent Docker image locally, run the following command:

```bash
./build/build-agent.sh
```

**Client Binary**
To build the CLI tool binary locally, run the following command:

```bash
./build/build-client.sh
```

## Contributing

Contributions are welcome! Before contributing, please read our Contributing Guidelines.

## License

This project is licensed under the MIT License.
