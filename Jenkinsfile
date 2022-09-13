pipeline{
    agent any
    parameters{
     booleanParam(
         name: 'APPROVALPASS',
         defaultValue: false,
         description: '同意部署吗？'
     )
     string(
       description:'输入docker版本',
       name:'DOCKERVERSION',
       defaultValue: '',
     )
     text(
       name: 'KUBRCHONFIG',
       defaultValue: '--kubeconfig=/root/.kube/config',
       description: 'kebeconfig'
     )
   }
    environment {
      REGISTRY_ENDPOINT = "https://registry.cn-hangzhou.aliyuncs.com/v2/"
      REGISTRY_CERTS = "60f91bd9-f443-4c96-948f-799d856b9f9e"
    }

    stages{
        stage("composer"){
            steps{
                echo "====++++composer++++===="
                sh '/usr/local/go/bin/go build -o ./savevault cmd/cmd.go'
            }
        }
    }
    post{
        always{
            echo "========always========"
        }
        success{
            echo "========pipeline executed successfully ========"
        }
        failure{
            echo "========pipeline execution failed========"
        }
    }
}