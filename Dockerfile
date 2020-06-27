ARG SERVER_NAME

ARG VIRTUAL_REPO_NAME

FROM ${SERVER_NAME}-${VIRTUAL_REPO_NAME}.jfrog.io/alpine:3.11.5

CMD printf "\nCONGRATULATIONS!!!\n\nYou have just set up your first Docker repository with the new JFrog Platform!\n\n"

