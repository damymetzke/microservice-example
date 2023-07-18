# Use the official Python base image
FROM python:3.11

# Set the working directory inside the container
WORKDIR /app

# Copy the PDM project files to the container
COPY . /app

# Install PDM
RUN pip install pdm

# Install project dependencies using PDM
RUN pdm install --no-lock

# Expose any necessary ports (replace 8000 with the actual port number)
EXPOSE 80

# Set the default command to run the application
CMD [ "pdm", "run", "src/main.py" ]
