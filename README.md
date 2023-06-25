# BoarScout Backend

Welcome to the BoarScout Backend, a Go Fiber application that powers the BoarScout app available on the app store. By open sourcing this repository, we offer the following benefits:

- Simplify the process of implementing new features in your team's scouting app.
- Ensure transparency for your data while using BoarScout.
- Enable automatic database sheet addition for your team.

## :warning: NOTE
BoarScout is undergoing a rewrite this summer. Expect new features and functionality. The contents of this repository may be frequently updated during this time.

## Getting started

Follow these steps to get started:

1. Install Go from [https://go.dev/dl/](https://go.dev/dl/). Choose and install the version appropriate for your operating system.
2. Install all required packages by running the following command in your terminal:

    ```shell
    go get -u all
    ```

3. Replace the `creds.json.txt` file with your Google account secret credentials and change the file extension from `.txt` to `.json`.
4. Create a `.env` file with the following variables:
    - `API_URL`: Enter a Discord webhook to receive logs whenever the app is opened.
    - `BL_API_KEY`: Include the API key for the Blue Alliance API.
    - `ACCESS_CODE`: Use the code your team will use to log into BoarScout.

5. Run the application with the following command:

    ```shell
    go run main.go
    ```
