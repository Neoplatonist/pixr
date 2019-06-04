# Pixr: Simplistic File Sharing

**Objective**:

Create a simplistic file sharing site that allows users to quickly upload content without the need for an account. This web app should be crafted in such a way that new features can be added with ease as the project matures.


**Languages**:
- Frontend
  - HTML5, CSS3, & Javascript ES6
- Backend
  - Golang

**Frameworks**:
- Frontend
  - [Svelte](https://github.com/sveltejs/svelte)
- Backend
  - [Labstack's Echo](http://github.com/labstack/echo)

----
## Product Requirements

### Pages
- [ ] Home Page
- [ ] Upload Page
- [ ] About Page
- [ ] Terms of Service
- [ ] Privacy Policy

### User Stories
#### Home Page
1. [ ] The user can sort images by date.
2. [ ] The user can keep seeing more images the more they scroll (infinite scroll).
3. [ ] The user can click an image to get a larger view in a modal.

#### Upload Page
1. [ ] The user can use a button to select files from their computer.
2. [ ] The user can drag and drop a selection of files.
3. [ ] The user can upload photos in bulk.
4. [ ] The user can remove photos from the upload preview before uploading.
5. [ ] The user is only able to upload image files.
6. [ ] The user can only upload a max of 10 images in bulk.
7. [ ] The user may need to use reCaptcha to verify they are not a bot when uploading.

#### About Page
1. [ ] The user can read more about the site and its rules.

----

## Backend Requirements
- Config File
- CLI Flags
- Limit to 10 files
- Rate Limit based on user ip addresses
- Verify uploader via reCaptcha
- Attempt to store all files uploaded
- Return single errors for individual files not uploaded
- Files:
  - Root Directory "data/"
  - Separated by date into folders "data/year-month-day/"
  - File named as date(unix styled)_ipaddress_bulkiterator.extension
  - Example: "data/2019-06-03/1559489386_localhost_0.jpg"
- Database:
  - Use MySQL
  - Table:
    - ID
    - Name
    - IP Address
    - File Location
    - Created At
  - API:
    - Upload
    - List of Images
    - Single Image
