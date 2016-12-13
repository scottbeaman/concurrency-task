# Concurrency context task
Currenty this program fires off requests concurrently to urls and when the first one comes back it cancels the rest of the requests using context.

## To run
```bash
    go run main.go
```

## Lesson 3
- Change the program not to cancel the requests but just to output the time taken for each request and print to the console.
