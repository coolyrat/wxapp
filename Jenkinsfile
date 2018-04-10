node {
    checkout scm;
    def wxappImage = docker.build("demo/wxapp");
    docker.withRegistry('http://harbor.codework.tech:8090', 'jenkins-harbor') {
        wxappImage.push("latest");
    }
    sh 'echo "push success"';
}