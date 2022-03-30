pipeline {
  agent any
  stages {
    stage('run tekton'){
        steps{
            tektonCreateRaw input: 'deploy/test/pipeline.yaml', inputType: 'FILE', namespace: 'tekton-devops-pipeline'
        }
    }
  }
}