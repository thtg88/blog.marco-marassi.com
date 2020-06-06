---
title: 'Run a serverless Laravel app with queue workers on AWS Lambda using Bref'
date: '2020-06-07'
---

Have you ever wanted to leverage the power of serverless computing but always felt stuck by AWS Lambda not providing native support for PHP? Read along!

I started messing around with AWS Lambda at my current job for a NodeJS project, and I wondered why, such a popular scripting language like PHP, which seems to be made exactly for the job of a Lambda function, is not supported natively.

Luckily someone has thought this before me, and setup a project called [Bref](https://bref.sh/).

Bref creates PHP bindings for AWS Lambda and abstracts away the complexity of Lambda layers, while also leveraging the [Serverless framework](https://www.serverless.com/), for pre-packaging the application and uploading it to your AWS account.

If you are wondering how is that going to integrate with a framework like Laravel, the guys at Bref have thought about that and provided a [handy guide](https://bref.sh/docs/frameworks/laravel.html).

Unfortunately there isn't quite yet a "native" support for running queue workers like `queue:work`, for example to send emails asynchronously from your application, as the Lambda function is a short-lived function which terminates once the execution finishes, Lambda is not really made for long-lived workers processes.

But I still like the idea of having my whole application in a single service (Lambda) without having to summon EC2, Fargate, or ECS, which can get definitely more expensive than Lambda.

We won't get defeated that easily! There are ways around it, here's how.

### Step-by-step Guide

1. Install the npm `serverless` package globally:

```
npm i -g serverless
```

2. Install the AWS CLI on your development machine, on macOS you can use Brew:

```
brew install awscli
```

For other operating systems, make sure to check the [AWS CLI documentation](https://docs.aws.amazon.com/cli/latest/userguide/install-cliv2.html).

3. Login to your AWS account, and create a user with programmatic access, for the Serverless framework to be able to access and upload resources.
While you can assign an `AdministratorAccess` access policy to the user for testing purposes, I'd recommend to use a stricter policy that just gives the required permissions to Serverless.
Their documentation [recommends](https://www.serverless.com/framework/docs/providers/aws/guide/credentials#creating-aws-access-keys) to create a policy with the following configuration:

```
{
    "Statement": [{
        "Action": [
            "apigateway:*",
            "cloudformation:CancelUpdateStack",
            "cloudformation:ContinueUpdateRollback",
            "cloudformation:CreateChangeSet",
            "cloudformation:CreateStack",
            "cloudformation:CreateUploadBucket",
            "cloudformation:DeleteStack",
            "cloudformation:Describe*",
            "cloudformation:EstimateTemplateCost",
            "cloudformation:ExecuteChangeSet",
            "cloudformation:Get*",
            "cloudformation:List*",
            "cloudformation:UpdateStack",
            "cloudformation:UpdateTerminationProtection",
            "cloudformation:ValidateTemplate",
            "dynamodb:CreateTable",
            "dynamodb:DeleteTable",
            "dynamodb:DescribeTable",
            "dynamodb:DescribeTimeToLive",
            "dynamodb:UpdateTimeToLive",
            "ec2:AttachInternetGateway",
            "ec2:AuthorizeSecurityGroupIngress",
            "ec2:CreateInternetGateway",
            "ec2:CreateNetworkAcl",
            "ec2:CreateNetworkAclEntry",
            "ec2:CreateRouteTable",
            "ec2:CreateSecurityGroup",
            "ec2:CreateSubnet",
            "ec2:CreateTags",
            "ec2:CreateVpc",
            "ec2:DeleteInternetGateway",
            "ec2:DeleteNetworkAcl",
            "ec2:DeleteNetworkAclEntry",
            "ec2:DeleteRouteTable",
            "ec2:DeleteSecurityGroup",
            "ec2:DeleteSubnet",
            "ec2:DeleteVpc",
            "ec2:Describe*",
            "ec2:DetachInternetGateway",
            "ec2:ModifyVpcAttribute",
            "events:DeleteRule",
            "events:DescribeRule",
            "events:ListRuleNamesByTarget",
            "events:ListRules",
            "events:ListTargetsByRule",
            "events:PutRule",
            "events:PutTargets",
            "events:RemoveTargets",
            "iam:AttachRolePolicy",
            "iam:CreateRole",
            "iam:DeleteRole",
            "iam:DeleteRolePolicy",
            "iam:DetachRolePolicy",
            "iam:GetRole",
            "iam:PassRole",
            "iam:PutRolePolicy",
            "iot:CreateTopicRule",
            "iot:DeleteTopicRule",
            "iot:DisableTopicRule",
            "iot:EnableTopicRule",
            "iot:ReplaceTopicRule",
            "kinesis:CreateStream",
            "kinesis:DeleteStream",
            "kinesis:DescribeStream",
            "lambda:*",
            "logs:CreateLogGroup",
            "logs:DeleteLogGroup",
            "logs:DescribeLogGroups",
            "logs:DescribeLogStreams",
            "logs:FilterLogEvents",
            "logs:GetLogEvents",
            "logs:PutSubscriptionFilter",
            "s3:CreateBucket",
            "s3:DeleteBucket",
            "s3:DeleteBucketPolicy",
            "s3:DeleteObject",
            "s3:DeleteObjectVersion",
            "s3:GetObject",
            "s3:GetObjectVersion",
            "s3:ListAllMyBuckets",
            "s3:ListBucket",
            "s3:PutBucketNotification",
            "s3:PutBucketPolicy",
            "s3:PutBucketTagging",
            "s3:PutBucketWebsite",
            "s3:PutEncryptionConfiguration",
            "s3:PutObject",
            "sns:CreateTopic",
            "sns:DeleteTopic",
            "sns:GetSubscriptionAttributes",
            "sns:GetTopicAttributes",
            "sns:ListSubscriptions",
            "sns:ListSubscriptionsByTopic",
            "sns:ListTopics",
            "sns:SetSubscriptionAttributes",
            "sns:SetTopicAttributes",
            "sns:Subscribe",
            "sns:Unsubscribe",
            "states:CreateStateMachine",
            "states:DeleteStateMachine"
        ],
        "Effect": "Allow",
        "Resource": "*"
    }],
    "Version": "2012-10-17"
}
```

You can find the above on a [GitHub Gist](https://gist.github.com/ServerlessBot/7618156b8671840a539f405dea2704c8).

Once the user is created with the attached policy, copy their access key ID and secret key ID.

4. Configure your AWS CLI to use the credentials created above with:

```
aws configure
```

You will be prompted to insert the above user's access key ID, secret key ID, region where you plan to place your assets and Lambda, and default output format. Feel free to leave the last one empty.

5. Now it's time to finally move to our project! Start by downloading a few packages:

```
composer require aws/aws-sdk-php bref/bref christoph-kluge/bref-sqs-laravel
```

What is going on here? We are downloading the PHP SDK for AWS to be able to put stuff on the SQS queue, Bref for our Lambda Serverless bindings, and [Bref SQS Laravel](https://github.com/christoph-kluge/bref-sqs-laravel) to have a worker run on Lambda.

6. Create a `serverless.yml` in your Laravel project's root folder, and copy the following content in it:

```
service: your-app-name

provider:
    name: aws
    region: us-west-1 # Make sure this matches the region of your SQS queue and the region you set when you did `aws configure`
    runtime: provided
    environment:
        APP_DEBUG: false
        APP_ENVIRONMENT: production
        # Logging to stderr allows the logs to end up in Cloudwatch
        LOG_CHANNEL: stderr
        # We cannot store sessions to disk: if you don't need sessions (e.g. API) then use `array`
        # If you write a website, use `cookie` or store sessions in database.
        SESSION_DRIVER: array
        VIEW_COMPILED_PATH: /tmp/storage/framework/views

plugins:
    - ./vendor/bref/bref

package:
  exclude:
    - node_modules/**
    - public/storage
    - resources/assets/**
    - storage/**
    - tests/**

functions:
    website:
        handler: public/index.php
        timeout: 28 # in seconds (API Gateway has a timeout of 29 seconds)
        layers:
            - ${bref:layer.php-74-fpm}
        events:
            -   http: 'ANY /'
            -   http: 'ANY /{proxy+}'
    artisan:
        handler: artisan
        timeout: 120 # in seconds
        layers:
            - ${bref:layer.php-74} # PHP
            - ${bref:layer.console} # The "console" layer
    queue:
        handler: artisan-lambda
        environment:
            ARTISAN_COMMAND: 'sqs:work sqs --tries=3 --sleep=1 --delay=1'
        layers:
            - ${bref:layer.php-74}
        events:
            - sqs:
                  arn: arn:aws:sqs:region:XXXXXX:default-queue # REPLACE THIS
                  batchSize: 10
```

Feel free to add any additional environment variables in the `environment` section, that you need to be passed to your application, that differ from the ones in your dev `.env` file.

**Be careful not to check into version control any secrets like API keys or password!**

Also be aware that some environment variables are reserved for AWS Lambda as they are passed by their environment. These are:

- `_HANDLER`: The handler location configured on the function.
- `AWS_REGION`: The AWS Region where the Lambda function is executed.
- `AWS_EXECUTION_ENV`: The runtime identifier, prefixed by AWS_Lambda_—for example, AWS_Lambda_java8.
- `AWS_LAMBDA_FUNCTION_NAME`: The name of the function.
- `AWS_LAMBDA_FUNCTION_MEMORY_SIZE`: The amount of memory available to the function in MB.
- `AWS_LAMBDA_FUNCTION_VERSION`: The version of the function being executed.
- `AWS_LAMBDA_LOG_GROUP_NAME`, `AWS_LAMBDA_LOG_STREAM_NAME`: The name of the Amazon CloudWatch Logs group and stream for the function.
- `AWS_ACCESS_KEY_ID`, `AWS_SECRET_ACCESS_KEY`, `AWS_SESSION_TOKEN`: The access keys obtained from the function's execution role.
- `AWS_LAMBDA_RUNTIME_API`: (Custom runtime) The host and port of the runtime API.
- `LAMBDA_TASK_ROOT`: The path to your Lambda function code.
- `LAMBDA_RUNTIME_DIR`: The path to runtime libraries.
- `TZ`: The environment's time zone (UTC). The execution environment uses NTP to synchronize the system clock.

See the [AWS docs](https://docs.aws.amazon.com/lambda/latest/dg/configuration-envvars.html#configuration-envvars-runtime) for an updated list.

7. Create an `artisan-lambda` file in your project's root directory and paste in the following content:

```
#!/opt/bin/php
<?php declare(strict_types=1);

use App\Console\Kernel;
use Illuminate\Contracts\Console\Kernel as BaseKernel;
use Symfony\Component\Console\Input\StringInput;
use Symfony\Component\Console\Output\ConsoleOutput;

$appRoot = getenv('LAMBDA_TASK_ROOT');
require_once $appRoot . '/vendor/autoload.php';
require_once $appRoot . '/bootstrap/app.php';

/** @var Kernel $kernel */
$kernel = app(BaseKernel::class);
$kernel->bootstrap();

$status = $kernel->handle(
    $input = new StringInput(getenv('ARTISAN_COMMAND')),
    new ConsoleOutput
);

$kernel->terminate($input, $status);
```

8. Make sure to edit `app/Providers/AppServiceProvider.php` so that directory is present (Laravel does not create it automatically):

```
public function boot()
{
    // Make sure the directory for compiled views exist
    if (! is_dir(config('view.compiled'))) {
        mkdir(config('view.compiled'), 0755, true);
    }
}
```

9. Configure your SQS queue. Assuming you already have created a queue in your AWS SQS dashboard, make sure to create a programmatic access user with SQS full access, download their credentials, and add them to your `.env` file:

```
QUEUE_CONNECTION=sqs
AWS_SQS_ACCESS_KEY_ID=changeme
AWS_SQS_SECRET_ACCESS_KEY=changeme
SQS_PREFIX=https://sqs.us-east-1.amazonaws.com/your-account-id
SQS_QUEUE=your-queue-name
AWS_SQS_DEFAULT_REGION=your-region-name
```

And change your `config/queue.php` in the `sqs` driver section:

```
'sqs' => [
    'driver' => 'sqs',
    'key' => env('AWS_SQS_ACCESS_KEY_ID'),
    'secret' => env('AWS_SQS_SECRET_ACCESS_KEY'),
    'prefix' => env('SQS_PREFIX', 'https://sqs.us-east-1.amazonaws.com/your-account-id'),
    'queue' => env('SQS_QUEUE', 'your-queue-name'),
    'suffix' => env('SQS_SUFFIX'),
    'region' => env('AWS_SQS_DEFAULT_REGION', 'us-east-1'),
],
```

We change the default AWS key env var names, as the original ones are reserved by AWS Lambda and would have conflicted with ours.

10. Now it's finally time to deploy your application! Run:

```
serverless deploy
```

This will create all the necessary AWS resources for you, and should work straight out of the box.

Happy coding! :)

### Further reading

- [Troubleshooting your Laravel Bref application on Lambda](https://bref.sh/docs/frameworks/laravel.html#troubleshooting)
- [Document how to use SQS to trigger Lambda](https://github.com/brefphp/bref/issues/421)
