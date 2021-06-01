pipeline {   
    agent any   
    
    environment {       
    registry = "btomasd/k8scicd"       
    GOCACHE = "/tmp"   
    }   
    stages {       
        stage('Build the deployment exec') {           
        agent {               
        docker {                   
            image 'golang'              
            }           
        }           
        steps {               
            // Create our project directory.               
                sh 'cd ${GOPATH}/src'               
                sh 'mkdir -p ${GOPATH}/src/hello-world'               
                // Copy all files in our Jenkins workspace to our project directory.                              
                sh 'cp -r ${WORKSPACE}/* ${GOPATH}/src/hello-world'
                sh 'ls'                
                // Build the app.
                //sh 'go mod init localhost/main'         
                sh 'go build'                        
                 }           
        }       
            stage('Test the Executable') {           
                agent {               
                    docker {                   
                    image 'golang'              
                    }          
                    }           
                steps {                               // Create our project directory.               
                        sh 'cd ${GOPATH}/src'               
                        sh 'mkdir -p ${GOPATH}/src/hello-world'               // Copy all files in our Jenkins workspace to our project directory.                              
                        sh 'cp -r ${WORKSPACE}/* ${GOPATH}/src/hello-world'               
                        // Remove cached test results.               
                        sh 'go clean -cache'               
                        // Run Unit Tests.               
                        sh 'go test ./... -v -short'                      
                        }       
                        
                    }       
        stage('DDockerize') {           
            environment {               
                registryCredential = 'dockerHub'           
                }           
                steps{               
                    script {                   
                        def appimage = docker.build registry + ":$BUILD_NUMBER"                   
                        docker.withRegistry( '', registryCredential ) 
                        { 
                                                appimage.push()                       
                                                appimage.push('latest')                  
                            }              
                            }  
                            } 
                                    
                    }       
        // stage ('Deploy') {           
        //     steps {               
        //         script{                   
        //         def image_id = registry + ":$BUILD_NUMBER"                   
        //         sh "ansible-playbook playbook.yml --extra-vars \"image_id=${image_id}\""              
        //         }          
        //         }       
                //  }   
                }
}
