def build() {
    echo "building the application..."
}

def test(){
    echo "testing the application..."
}

def deploy(){
    echo "deploying the application..."
}

// return all defined functions so they are accessible in the jenkinsfile
return this
