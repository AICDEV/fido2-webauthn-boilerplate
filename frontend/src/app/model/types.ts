import { JwtPayload } from 'jwt-decode';


export interface ILoginData {
  id: string;
  rawId: string;
  response: {
    authenticatorData: string;
    clientDataJSON: string;
    signature: string;
    userHandle: string;
  };
  type: 'public-key';
}

export interface WrappedPublicKeyCredentialRequestOptionsJSON {
  publicKey: PublicKeyCredentialRequestOptionsJSON;
}

export interface PublicKeyCredentialRequestOptionsJSON {
  challenge: string;
  timeout?: number;
  rpId?: string;
  allowCredentials?: PublicKeyCredentialDescriptorJSON[];
  userVerification?: UserVerificationRequirement;
  extensions?: AuthenticationExtensionsClientInputs;
}

export interface PublicKeyCredentialDescriptorJSON {
  id: string;
  type: PublicKeyCredentialType;
  transports?: AuthenticatorTransport[];
}


export interface AuthenticationCredential extends PublicKeyCredentialFuture {
  response: AuthenticatorAssertionResponse;
}

export interface AuthenticatorAttestationResponseJSON {
  clientDataJSON: string;
  attestationObject: string;
  authenticatorData?: string;
  transports?: AuthenticatorTransport[];
  publicKeyAlgorithm?: COSEAlgorithmIdentifier;
  publicKey?: string;
}

export interface RegistrationResponseJSON {
  id: string;
  rawId: string;
  response: AuthenticatorAttestationResponseJSON;
  authenticatorAttachment?: AuthenticatorAttachment;
  clientExtensionResults: AuthenticationExtensionsClientOutputs;
  type: PublicKeyCredentialType;
}

export interface AuthenticatorAssertionResponseJSON {
  clientDataJSON: string;
  authenticatorData: string;
  signature: string;
  userHandle?: string;
}

export interface AuthenticationResponseJSON {
  id: string;
  rawId: string;
  response: AuthenticatorAssertionResponseJSON;
  authenticatorAttachment?: AuthenticatorAttachment;
  clientExtensionResults: AuthenticationExtensionsClientOutputs;
  type: PublicKeyCredentialType;
}

export type PublicKeyCredentialJSON = RegistrationResponseJSON | AuthenticationResponseJSON;

export interface PublicKeyCredentialFuture extends PublicKeyCredential {
  type: PublicKeyCredentialType;
  isConditionalMediationAvailable?(): Promise<boolean>;
  parseCreationOptionsFromJSON?(
    options: PublicKeyCredentialCreationOptionsJSON,
  ): PublicKeyCredentialCreationOptions;
  parseRequestOptionsFromJSON?(
    options: PublicKeyCredentialRequestOptionsJSON,
  ): PublicKeyCredentialRequestOptions;
  toJSON?(): PublicKeyCredentialJSON;
}



export interface PublicKeyCredentialCreationOptionsJSON {
  rp: PublicKeyCredentialRpEntity;
  user: PublicKeyCredentialUserEntityJSON;
  challenge: string;
  pubKeyCredParams: PublicKeyCredentialParameters[];
  timeout?: number;
  excludeCredentials?: PublicKeyCredentialDescriptorJSON[];
  authenticatorSelection?: AuthenticatorSelectionCriteria;
  attestation?: AttestationConveyancePreference;
  extensions?: AuthenticationExtensionsClientInputs;
}

export interface PublicKeyCredentialUserEntityJSON {
  id: string;
  name: string;
  displayName: string;
}

export interface RegistrationCredential extends PublicKeyCredentialFuture {
  response: AuthenticatorAttestationResponse;
}

export interface IFinalizeLoginResponse {
  token: string;
}

export interface IToken extends JwtPayload {
  firstName: string;
  lastName: string;
  mail: string;
}


export enum AUTH_STATE {
  'VERIFIED' =  'VERIFIED',
  'UNKNOWN_USER' = 'UNKNOWN_USER',
  'NOT_VERIFIED' = 'NOT_VERIFIED',
  'HTTP_ERROR' = 'HTTP_ERROR',
  'AUTHENTICATOR_ERROR' = 'AUTHENTICATOR_ERROR',
}


export interface IRegisterResponse {
  publicKey: {
    rp: {
      name: string;
      id: string;
    },
    user: {
      name: string;
      displayName: string;
      id: string;
    },
    challenge: string;
    pubKeyCredParams: { type: 'public-key'; alg: number; }[];
    timeout: number;
    authenticatorSelection: {
      requireResidentKey: boolean;
      userVerification: 'preferred' | 'discouraged' | 'required';
    },
  },
}
