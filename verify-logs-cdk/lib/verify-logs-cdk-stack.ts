import * as cdk from "@aws-cdk/core";
import * as ecs from "@aws-cdk/aws-ecs";
import * as ec2 from "@aws-cdk/aws-ec2";
import * as iam from "@aws-cdk/aws-iam";

export class VerifyLogsCdkStack extends cdk.Stack {
  constructor(scope: cdk.Construct, id: string, props?: cdk.StackProps) {
    super(scope, id, props);
    const vpc = ec2.Vpc.fromLookup(this, "VPC", {
      vpcName: this.node.tryGetContext("vpcName"),
    });
    const cluster = new ecs.Cluster(this, "testCluster", {
      vpc: vpc,
    });


    const taskRole = new iam.Role(this, 'ECSTaskRole', {
      assumedBy: new iam.ServicePrincipal('ecs-tasks.amazonaws.com'),
    })
    taskRole.addManagedPolicy(iam.ManagedPolicy.fromAwsManagedPolicyName('CloudWatchLogsReadOnlyAccess'))
    // create a task definition with CloudWatch Logs
    const logging = new ecs.AwsLogDriver({
      streamPrefix: "/aws/ecs/verify-logs",
    });

    const taskDef = new ecs.FargateTaskDefinition(this, "MyTaskDefinition", {
      memoryLimitMiB: 512,
      cpu: 256,
      taskRole: taskRole,
    });

    taskDef.addContainer("AppContainer", {
      image: ecs.ContainerImage.fromRegistry("masahide/verify-logs"),
      logging: logging,
      environment: {
        GROUPNAME: "/aws/ecs/verify-logs",
        AWS_REGION: "ap-northeast-1",
      },
    });

    // Instantiate ECS Service with just cluster and image
    new ecs.FargateService(this, "FargateService", {
      cluster,
      taskDefinition: taskDef,
    });

    // The code that defines your stack goes here
  }
}
