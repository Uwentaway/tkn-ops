pipeline {
  agent any
  stages {
    stage('run tekton'){
        steps{
            tektonCreateRaw input: 'deploy/test/pipeline.yaml', inputType: 'FILE', namespace: 'tekton-devops-pipeline'
        }
    }
    stage('Static Analysis') {
    steps {
        recordIssues(
            tool: pyLint(pattern: '**/pylint.out'),
            unstableTotalAll: 100,
        )
    }
  }
}