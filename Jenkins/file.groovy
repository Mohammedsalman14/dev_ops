def build() {
    echo "Building the application..."
}

def test() {
    echo "Testing the application..."
}

def deploy() {
    echo "Deploying the application..."
}

// Expose the methods to Jenkinsfile
return this
