***Blog Api connecting to Postgres using Gin and Gorm***

**Instructions**

1. **Migrate Database:** 
    ```
    go run migrate/migrate.go
    ```

2. **Run Application:**
    ```
    go run main.go
    ```

3. **Create Post:**
    - **Endpoint:** localhost:3000/posts
    - **Method:** POST
    - **Request Body:**
        ```json
        {
            "title": "Greatness",
            "body": "Great people are usually impactful"
        }
        ```
    - **Response:**
        ```json
        {
            "post": {
                "ID": 3,
                "CreatedAt": "2024-03-03T16:26:48.875725Z",
                "UpdatedAt": "2024-03-03T16:26:48.875725Z",
                "DeletedAt": null,
                "Title": "Greatness",
                "Body": "Great people are usually impactful"
            }
        }
        ```

4. **Get All Posts:**
    - **Endpoint:** localhost:3000/posts
    - **Method:** GET
    - **Response:**
        ```json
        {
            "posts": [
                {
                    "ID": 2,
                    "CreatedAt": "2024-03-03T16:26:41.926829Z",
                    "UpdatedAt": "2024-03-03T17:55:54.59172Z",
                    "DeletedAt": null,
                    "Title": "Greatness",
                    "Body": "Great people are usually impactful"
                },
                {
                    "ID": 3,
                    "CreatedAt": "2024-03-03T16:26:48.875725Z",
                    "UpdatedAt": "2024-03-03T16:26:48.875725Z",
                    "DeletedAt": null,
                    "Title": "from req body",
                    "Body": "from req body"
                }
            ]
        }
        ```

5. **Get Post by ID:**
    - **Endpoint:** localhost:3000/posts/2
    - **Method:** GET
    - **Response:**
        ```json
        {
            "post": {
                "ID": 2,
                "CreatedAt": "2024-03-03T16:26:41.926829Z",
                "UpdatedAt": "2024-03-03T16:26:41.926829Z",
                "DeletedAt": null,
                "Title": "from req body",
                "Body": "from req body"
            }
        }
        ```

6. **Update Post by ID:**
    - **Endpoint:** localhost:3000/posts/2
    - **Method:** UPDATE
    - **Request Body:**
        ```json
        {
            "title": "Greatness",
            "body": "Great people are usually impactful"
        }
        ```
    - **Response:**
        ```json
        {
            "post": {
                "ID": 2,
                "CreatedAt": "2024-03-03T16:26:41.926829Z",
                "UpdatedAt": "2024-03-03T17:55:54.59172Z",
                "DeletedAt": null,
                "Title": "Greatness",
                "Body": "Great people are usually impactful"
            }
        }
        ```

7. **Delete Post by ID:**
    - **Endpoint:** localhost:3000/posts/2
    - **Method:** DELETE
