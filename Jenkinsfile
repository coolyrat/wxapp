node {
    checkout scm
    def wxappImage = docker.build("harbor.codework.tech:8090/demo/wxapp:latest")
    wxappImage.push('http://harbor.codework.tech:8090', 'jenkins-harbor')
    sh 'echo "push success"'
}