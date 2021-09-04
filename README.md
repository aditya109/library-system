# go-server

1. Create a mux server

   1. Create API: ✔

      | API Endpoint                   | Method Type | Function                                   |
      | ------------------------------ | ----------- | ------------------------------------------ |
      | `/api/v1/books`                | GET         | Get all books' details paginated           |
      | `/api/v1/bookByBookId`         | GET         | Gets 1 book's details via book id          |
      | `/api/v1/bookByBookName`       | GET         | Gets 1 book's details via book name        |
      | `/api/v1/bookByBookAuthorName` | GET         | Gets 1 book's details via book author name |
      | `/api/v1/bookByPrice`          | GET         | Gets 1 book's details                      |
      | `/api/v1/bookByIsbn`           | GET         | Gets 1 book's details                      |
      | `/api/v1/book`                 | POST        | Add 1 book to DB                           |
      | `/api/v1/books`                | POST        | Add multiple books to DB                   |

2. Connect this server to `mongo-db-container`. ✔

   ```shell
   docker run --name mongo-db -p 27017:27017 -d mongo:latest
   ```

3. Dockerize this app. ✔

4. Create orchestration manifests.
