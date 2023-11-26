## Prerequisites

Things you need to run this application:

- Docker

## Usage

1. Starting Docker Container  
   Use Docker Compose to set up the necessary environment.  
   Run the following command to start the container in the background.
    ```bash
    $ docker-compose up -d
    ```

2. Building the Application
    Next, use the Makefile to build the application. Run the following command.
    ```bash
    $ make build
    ```

3. Running the Application
    Once the build is complete, run the application with the following command.

    ```bash
    $ make run
    ```

## License