node {
    stage('Start testing environment') {
        echo "Workspace is : ${env.WORKSPACE}"
        sh 'docker-compose up -d'
    }

    stage('Stop testing environment') {
        sh 'docker-compose down'
    }
}