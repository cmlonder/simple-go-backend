node {
    stage('clean') {
        cleanWs()
    }

     stage('checkout') {
        checkout scm
     }

    stage('Start testing environment') {
        echo "Workspace is : ${env.WORKSPACE}"
        sh 'docker-compose up -d'
    }

    stage("run test") {
        sh 'java --version'
        bzt "load_test.jmx"
    }

    stage('Stop testing environment') {
        sh 'docker-compose down'
    }
}