# Firebase Storage Uploader
A small tool to upload your file to your Firebase Storage.

## Build From Source
### Prerequisites
1. latest Go
2. Git installed
3. Service Account Key from Firebase, Generate new private key from here https://console.firebase.google.com/project/[YOUR_FIREBASE_PROJECT]/settings/serviceaccounts/adminsdk
4. File you want to upload, e.g. `somefile.jpg`

### How to Build
1. Clone this repo
    ```sh
    $ git clone git@github.com:kodebineri/storage-uploader.git
    ```

2. Install dependencies and build
    ```sh
    $ cd storage-uploader
    $ go install
    $ go build
    ```

3. Set environment variable
    ```sh
    $ export STORAGE_BUCKET=something.appspot.com
    $ export CREDENTIAL_PATH=/your/service/account/key/json/path
    ```

4. Start uploading
    ```sh
    $ ./storage-uploader somefile.jpg
    ```

## Download Executables
Do not have time to build? Make sure you have already set your environment variables and download executable for your platform [here](https://drive.google.com/drive/folders/1yCnA6u1MmjKPKd_AqLyjcNFgJwLHj1wZ?usp=sharing) and execute, e.g.
```sh
$ storage-uploader_win_x64.exe somefile.jpg
```