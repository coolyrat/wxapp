node {
    checkout scm;
    docker.withRegistry('https://harbor.codework.tech:8090', 'jenkins-harbor') {
        def wxappImage = docker.build("harbor.codework.tech:8090/demo/wxapp");
        wxappImage.push("latest");
    }
    sh 'echo "push success"';
}