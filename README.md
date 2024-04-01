# Canvas for Backend Technical Test at Scalingo

## Instructions

* From this canvas, respond to the project which has been communicated to you by our team
* Feel free to change everything

## Project architecture

own
Copy code
````markdown
                    +-------------------+
                    | generateSource    |
                    |     Channel       |
                    +---------+---------+
                              |
         +----------------------------------------+
         |                                        |
         v                                        v
+------------------+                     +------------------+
| workerLanguages  |                     |   workerLicense  |
|    Worker Pool   |                     |   Worker Pool    |
+------------------+                     +------------------+
         |                     |                   |
         +---------------------|-------------------+
                               |
                    +----------v----------+
                    |        merge        |
                    +---------------------+
````

Description:

- **generateSource Channel**: This component generates task items and sends them to worker pools for processing.
- **workerLanguages Worker Pool** and **workerLicense Worker Pool**: These are pools of worker goroutines responsible for concurrently fetching languages and license information for each repository.
- **merge function**: Collects and merges the results from the worker pools.

This architecture follows the fan-out, fan-in concurrency pattern, where tasks are fanned out to multiple worker pools for parallel processing, and results are fanned in and merged.
## Execution

```
docker compose up
```

Application will be then running on port `5000`

### Setting Up Environment Variables

To configure the application's environment variables, follow these steps:

1. **Locate the `.env.example` file:**
   - In the root directory of the project, you will find a file named `.env.example`.

2. **Create a new `.env` file:**
   - Duplicate `.env.example` and rename the duplicated file to `.env`.

3. **Replace dummy values with actual values:**
   - Open the newly created `.env` file in a text editor.
   - Replace the dummy values provided in `.env.example` with your actual environment variable values.
   - Ensure that the real values are securely managed and not exposed publicly.

4. **Save the `.env` file:**
   - Once you have replaced the dummy values with real values, save the `.env` file.

By following these steps, you will properly configure the application's environment variables with the required values.

## Test

```
$ curl localhost:5000/ping
{ "status": "pong" }
```

```
$ curl localhost:5000/repos?language=javascript
[
    {
        "name": "example-repo",
        "languages": {
            "JavaScript": 49,
            "Ruby": 134218
         },
        "license": {
            "key": "mit",
            "name": "MIT License",
            "node_id": "MDc6TGljZW5zZTEz",
            "spdx_id": "MIT",
            "url": "https://api.github.com/licenses/mit"
        }
        ...
    },
    ...
]

```
