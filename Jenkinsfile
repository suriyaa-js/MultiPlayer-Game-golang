pipeline {
    agent any

    stages {

        stage('Build') {
            steps {
                echo 'Building the application...'
                script{
                    def test =2+3 > 4? 'cool' : 'not cool'
                    echo test
                }
            }
        }

        stage('Test') {
            steps {
                echo 'Running tests...'
            }
        }

        stage('Deploy') {
            steps {
                echo 'Deploying the application...'
            }
        }
    }
}
