/* eslint-disable */
/**
 * This file was automatically generated by json-schema-to-typescript.
 * DO NOT MODIFY IT BY HAND. Instead, modify the source JSONSchema file,
 * and run json-schema-to-typescript to regenerate this file.
 */

/**
 * Meshplay Connections are managed and unmanaged resources that either through discovery or manual entry are tracked by Meshplay. Learn more at https://docs.meshplay.khulnasoft.com/concepts/logical/connections
 */
export interface HttpsSchemasMeshplayKhulnasoftComComponentJson {
  /**
   * ID
   */
  id?: string;
  /**
   * Connection Name
   */
  name?: string;
  /**
   * Credential ID
   */
  credential_id?: string;
  /**
   * Connection Type
   */
  type: string;
  /**
   * Connection Subtype
   */
  sub_type?: string;
  /**
   * Connection Kind
   */
  kind: string;
  metadata?: {
    [k: string]: unknown;
  };
  /**
   * Connection Status
   */
  status:
    | "discovered"
    | "registered"
    | "connected"
    | "ignored"
    | "maintenance"
    | "disconnected"
    | "deleted"
    | "not found";
  /**
   * A Universally Unique Identifier used to uniquely identify entites in Meshplay. The UUID core defintion is used across different schemas.
   */
  user_id?: string;
  created_at?: string;
  updated_at?: string;
  deleted_at?: string;
}
