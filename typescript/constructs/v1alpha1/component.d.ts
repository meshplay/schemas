/* eslint-disable */
/**
 * This file was automatically generated by json-schema-to-typescript.
 * DO NOT MODIFY IT BY HAND. Instead, modify the source JSONSchema file,
 * and run json-schema-to-typescript to regenerate this file.
 */

/**
 * Components are the atomic units for designing infrastructure. Learn more at https://docs.meshplay.io/concepts/components
 */
export interface ComponentDefinitionsJson {
  /**
   * API Version of the component.
   */
  apiVersion: string;
  /**
   * Kind of the component.
   */
  kind: string;
  /**
   * Metadata contains additional information associated with the component.
   */
  metadata?: {
    /**
     * Description of the component.
     */
    description?: string;
    /**
     * Meshplay manages components in accordance with their specific capabilities. This field explicitly identifies those capabilities largely by what actions a given component supports; e.g. metric-scrape, sub-interface, and so on. This field is extensible. ComponentDefinitions made define a broad array of capabilities, which are in-turn dynamically interpretted by Meshplay for full lifecycle management.
     */
    capabilities?: {
      [k: string]: unknown;
    };
    /**
     * Name of the component.
     */
    name: string;
    /**
     * Version of the component.
     */
    version: string;
    [k: string]: unknown;
  };
  model: HttpsSchemasMeshplayIoModelJson;
  /**
   * Display name of the component.
   */
  displayName?: string;
  /**
   * Format specifies the format used in the `schema` field. JSON will be used as a default format.
   */
  format?: "JSON" | "CUE";
  /**
   * JSON schema of the component.
   */
  schema: string;
}
/**
 * Model of the component. Learn more at https://docs.meshplay.io/concepts/models
 */
export interface HttpsSchemasMeshplayIoModelJson {
  /**
   * The name for the model.
   */
  name: string;
  /**
   * The display name for the model.
   */
  displayName?: string;
  /**
   * Status of model, e.g. Registered, Ignored, Enabled ...
   */
  status: string;
  /**
   * Version of the model.
   */
  version: string;
  /**
   * Category of the model.
   */
  category: string;
  /**
   * Sub-category of the model.
   */
  subCategory?: string;
}
