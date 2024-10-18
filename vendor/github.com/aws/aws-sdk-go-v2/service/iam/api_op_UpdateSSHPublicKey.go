// Code generated by smithy-go-codegen DO NOT EDIT.

package iam

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/iam/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Sets the status of an IAM user's SSH public key to active or inactive. SSH
// public keys that are inactive cannot be used for authentication. This operation
// can be used to disable a user's SSH public key as part of a key rotation work
// flow.
//
// The SSH public key affected by this operation is used only for authenticating
// the associated IAM user to an CodeCommit repository. For more information about
// using SSH keys to authenticate to an CodeCommit repository, see [Set up CodeCommit for SSH connections]in the
// CodeCommit User Guide.
//
// [Set up CodeCommit for SSH connections]: https://docs.aws.amazon.com/codecommit/latest/userguide/setting-up-credentials-ssh.html
func (c *Client) UpdateSSHPublicKey(ctx context.Context, params *UpdateSSHPublicKeyInput, optFns ...func(*Options)) (*UpdateSSHPublicKeyOutput, error) {
	if params == nil {
		params = &UpdateSSHPublicKeyInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateSSHPublicKey", params, optFns, c.addOperationUpdateSSHPublicKeyMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateSSHPublicKeyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateSSHPublicKeyInput struct {

	// The unique identifier for the SSH public key.
	//
	// This parameter allows (through its [regex pattern]) a string of characters that can consist of
	// any upper or lowercased letter or digit.
	//
	// [regex pattern]: http://wikipedia.org/wiki/regex
	//
	// This member is required.
	SSHPublicKeyId *string

	// The status to assign to the SSH public key. Active means that the key can be
	// used for authentication with an CodeCommit repository. Inactive means that the
	// key cannot be used.
	//
	// This member is required.
	Status types.StatusType

	// The name of the IAM user associated with the SSH public key.
	//
	// This parameter allows (through its [regex pattern]) a string of characters consisting of upper
	// and lowercase alphanumeric characters with no spaces. You can also include any
	// of the following characters: _+=,.@-
	//
	// [regex pattern]: http://wikipedia.org/wiki/regex
	//
	// This member is required.
	UserName *string

	noSmithyDocumentSerde
}

type UpdateSSHPublicKeyOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateSSHPublicKeyMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsquery_serializeOpUpdateSSHPublicKey{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpUpdateSSHPublicKey{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "UpdateSSHPublicKey"); err != nil {
		return fmt.Errorf("add protocol finalizers: %v", err)
	}

	if err = addlegacyEndpointContextSetter(stack, options); err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = addClientRequestID(stack); err != nil {
		return err
	}
	if err = addComputeContentLength(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = addComputePayloadSHA256(stack); err != nil {
		return err
	}
	if err = addRetry(stack, options); err != nil {
		return err
	}
	if err = addRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = addRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addSpanRetryLoop(stack, options); err != nil {
		return err
	}
	if err = addClientUserAgent(stack, options); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addSetLegacyContextSigningOptionsMiddleware(stack); err != nil {
		return err
	}
	if err = addTimeOffsetBuild(stack, c); err != nil {
		return err
	}
	if err = addUserAgentRetryMode(stack, options); err != nil {
		return err
	}
	if err = addOpUpdateSSHPublicKeyValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateSSHPublicKey(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRecursionDetection(stack); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	if err = addDisableHTTPSMiddleware(stack, options); err != nil {
		return err
	}
	if err = addSpanInitializeStart(stack); err != nil {
		return err
	}
	if err = addSpanInitializeEnd(stack); err != nil {
		return err
	}
	if err = addSpanBuildRequestStart(stack); err != nil {
		return err
	}
	if err = addSpanBuildRequestEnd(stack); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opUpdateSSHPublicKey(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "UpdateSSHPublicKey",
	}
}
