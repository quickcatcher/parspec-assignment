# parspec-assignment

This document provides instructions for setting up and running the parspec-assignemnet order backend system.

## Prerequisites

Before you begin, ensure that you have the following software installed:

*   **Go:**  Go version 1.18 or later is recommended. You can download and install Go from [https://go.dev/](https://go.dev/).  After installation, make sure to set up your Go workspace and add the Go binaries to your PATH environment variable.

*   **MySQL:** MySQL Server is required for database operations. You can download and install MySQL from [https://www.mysql.com/](https://www.mysql.com/). After installation, ensure that the MySQL server is running. You'll also need a MySQL client (like the `mysql` command-line tool or MySQL Workbench) to create the database and user.

## Database Setup

1.  **Create Database:**  Use your MySQL client to create the database that the application will use.

    ```sql
    CREATE DATABASE parspec;
    ```

2.  **Create User (Optional but Recommended):**  It's good security practice to create a dedicated MySQL user for your application. Replace `your_username` and `your_password` with your desired credentials:

    ```sql
    CREATE USER 'your_username'@'localhost' IDENTIFIED BY 'your_password';
    GRANT ALL PRIVILEGES ON your_database_name.* TO 'your_username'@'localhost';
    FLUSH PRIVILEGES;
    ```

3.  **Create Table:** Here's the SQL query used to create the table:

    ```sql
    CREATE TABLE IF NOT EXISTS orders (
    `order_id` INT PRIMARY KEY auto_increment,
    `user_id` INT,
    `item_ids` TEXT,
    `total_amount` DECIMAL(10,2),
    `status` ENUM('Pending', 'Processing', 'Completed') DEFAULT 'Pending',
    `processing_time` FLOAT,
    `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP ,
    `updated_at` DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
    );
    ```

## Setting up the Application

1.  **Clone the Repository:** Clone the project repository to your local machine.


2.  **Install Dependencies:** The project uses Go modules for dependency management. Navigate to the project directory in your terminal and run the following command to download the required dependencies:

    ```bash
    go mod tidy
    ```

3.  **.env File:** Create a `.env` file in the root directory of the project and add your MySQL database credentials:

    ```
    DB_USER=your_mysql_username
    DB_PASSWORD=your_mysql_password
    DB_NAME=parsepec
    DB_HOST=your_mysql_host  # e.g., localhost or 127.0.0.1
    DB_PORT=your_mysql_port  # e.g., 3306
    ```

## Running the Application

1.  **Start the Server:**  Navigate to the project directory in your terminal and run the following command to start the application:

    ```bash
    go run main.go
    ```

    The server should start listening on port 9000. You should see a message like "Server listening on port 9000" in your terminal.

2.  **Testing the API:** You can use tools like `curl`, Postman, or a similar HTTP client to test the API endpoints.  See the "Example API Requests and Responses" section below for examples.

## Example API Requests and Responses

### Create Order

```bash
curl --location 'http://localhost:9000/api/v1/order' \
--header 'Content-Type: application/json' \
--data '{
    "user_id": 1,
    "item_ids": "1,2,3,4",
    "total_amount": 1000.00
}'

Response
{
    "message": "Order created successfully",
    "model": {
        "order_id": 22
    }
}
```

### Get Order Status

```bash
curl --location 'http://localhost:9000/parspec/order/1'

Response
{
    "message": "Order Found",
    "model": {
        "status": "Pending",
        "item_ids": "1,2,3",
        "total_amount": 1000
    }
}
```

### Get Metrics

```bash
curl --location 'http://localhost:9000/parspec/metrics'

Response
{
    "total_orders_processed": 3,
    "average_processing_time": 5,
    "order_status_counts": {
        "Completed": 3,
        "Pending": 0,
        "Processing": 1
    }
}
```


**Directory Structure Explanation:**

*   **`/` (Root Directory):** Contains the main application file (`main.go`).
*   **`route/`:**  Contains the routing logic. `route.go` defines the API endpoints and maps them to handler functions (usually in the `controllers` package).
*   **`core/`:** Contains the core business logic and data models of the application.
    *   **`domain/`:**  Defines the domain models (entities) used in the application.  These are plain Go structs that represent the data.
    *   **`service/`:** Contains the business logic. Services operate on domain models and often interact with the persistence layer.
    *   **`persistence/`:** Handles the interaction with the database.  This layer uses the Beego ORM to perform CRUD (Create, Read, Update, Delete) operations on the database.
*   **`middleware/`:** Contains middleware functions. Middleware are functions that run before the request handlers and can be used for tasks like authentication, authorization, logging, database connection management, etc.

This structure promotes separation of concerns, making the code more organized, maintainable, and testable.  Each layer has a specific responsibility, and the dependencies between layers are clearly defined.

# Choosing an In-Memory Queue in Go: Channels and Alternatives

I used Go channels for the in-memory queue implementation in this project and have mentioned other possible options along with their trade-offs.

## Why Go Channels?

Go channels are a built-in feature of the Go language that provide a powerful and elegant way to manage concurrent communication and data transfer between goroutines.  For this specific project, Go channels were a suitable choice for the in-memory queue due to the following reasons:

*   **Simplicity:** Go channels are easy to use and understand.  They provide a clean and concise way to implement a queue without the need for external libraries or dependencies.  This simplicity is beneficial for keeping the project focused on the core concepts of order processing and asynchronous operations.

*   **Concurrency:** Go channels are designed for concurrent programming.  They allow goroutines to communicate and exchange data safely and efficiently.  This is essential for the asynchronous order processing requirement, where the HTTP server (running in one goroutine) needs to enqueue orders, and a background processing goroutine needs to dequeue and process them.

*   **Built-in:**  Channels are a core part of the Go language, so there's no need to add any external dependencies. This keeps the project lightweight and avoids potential dependency conflicts.

*   **Blocking Behavior:**  Channels have a built-in blocking behavior.  When a goroutine tries to receive from an empty channel, it will block until data is available.  This blocking behavior is exactly what's needed for a queue.  The order processing goroutine can wait on the channel until a new order is enqueued.

## Other Possible Options and Trade-offs

While Go channels are a good choice for this project, other options exist for implementing in-memory queues in Go. Here are some of them:

1.  **`sync/atomic` and Slices/Maps:** You could use `sync/atomic` operations along with slices or maps to create a concurrent-safe queue.

    *   **Trade-offs:** This approach is significantly more complex than using channels. You would need to manually manage locks and atomic counters to ensure data consistency. While possible, it adds a lot of boilerplate code and increases the risk of introducing concurrency bugs.  For this project, the added complexity is not justified.

2.  **Dedicated In-Memory Queue Libraries:** Several Go libraries provide in-memory queue implementations (e.g., `github.com/enrique-gonzalez/queue`).

    *   **Trade-offs:** These libraries often offer more advanced features than simple channels, such as persistence, priority queues, and more sophisticated queue management.  However, they introduce an external dependency. For this project, which focuses on core concepts, the added complexity of a library is not necessary.  If the project were to require more advanced queueing features, then a library might be a good choice.

3.  **Ring Buffer:** A ring buffer is another data structure that can be used to implement a queue.

    *   **Trade-offs:** Ring buffers can be very efficient, but they require careful management of indices and can be more complex to implement correctly than channels.  For this project, the performance benefits of a ring buffer are unlikely to be significant enough to justify the added complexity.

## Why Channels are the Best Fit for This Project

For this specific project, which focuses on demonstrating the fundamental principles of asynchronous order processing, Go channels provide the best balance of simplicity, efficiency, and suitability.  They are easy to use, well-suited for concurrent communication, and built into the Go language.  The other options, while potentially offering more advanced features or performance benefits, introduce unnecessary complexity for the scope of this project.  If the requirements of the project were to change (e.g., needing queue persistence, priority queues, or extremely high throughput), then other options might be considered. However, for a simple in-memory queue, channels are the most appropriate and idiomatic Go solution.
