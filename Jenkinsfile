pipeline {
    agent any

    environment {
        SONAR_HOST_URL = 'http://jenkinsIP:9000'
        SONAR_LOGIN = ''
        SONAR_PASSWORD = ''
    }

    stages {
        stage('Checkout') {
            steps {
                git url: 'http://gogIp:10880/username/intro-K8S-operator.git', branch: 'main'
            }
        }

        stage('SonarQube Analysis') {
            steps {
                withSonarQubeEnv('MyScanner') {
                    sh """
                        sonar-scanner \
                          -Dsonar.projectKey=intro-K8S-operator \
                          -Dsonar.sources=. \
                          -Dsonar.host.url=${SONAR_HOST_URL} \
                          -Dsonar.login=${SONAR_LOGIN} \
                          -Dsonar.password=${SONAR_PASSWORD}
                    """
                }
            }
        }
    }
}
