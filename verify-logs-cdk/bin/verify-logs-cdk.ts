#!/usr/bin/env node
import 'source-map-support/register';
import * as cdk from '@aws-cdk/core';
import { VerifyLogsCdkStack } from '../lib/verify-logs-cdk-stack';

const app = new cdk.App();
new VerifyLogsCdkStack(app, 'VerifyLogsCdkStack');
