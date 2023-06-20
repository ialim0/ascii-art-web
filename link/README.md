## create docker image
Open a terminal or command prompt, navigate to the directory where the Dockerfile is located (the root directory of your Go project), and run the following command:

"docker build -t my-go-app ."
This command wil help you create  an image name "my-go-app".The period (.) at the end of the command indicates that the Dockerfile is in the current directory.
# 
Once the image is built, you can run a Docker container using the following command:
"docker run -p 8080:8080 my-go-app"
##To view the labels of a Docker image, you can use the docker inspect command. For example, to inspect the labels of the "my-go-app" image, run the following command:
"docker inspect my-go-app"

# List running containers
"docker ps"
# Delete dangling images: Dangling images are the ones that are not associated with any tagged or named image. To remove dangling images, you can use the docker image prune command. This command will remove all unused images. 
"docker image prune"
# Clean up unused resources:
"docker system prune"
