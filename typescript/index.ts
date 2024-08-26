// This file is not automatically generated. It is created manually to export the constructs in the generated modules
// The generated modules are created in the dist folder and are named based on the constructs in the schema

// v1alpha1
import {
  CatalogData as CatalogDataT,
  CaveatsAndConsiderations,
} from "./constructs/v1alpha1/catalog_data";
import {
HttpsSchemasMeshplayCapabilityJson
} from "./constructs/v1alpha1/capability"

// v1alpha2
import { DesignSchema } from "./constructs/v1alpha2/design";
import { HttpsSchemasMeshplayRelationshipJson } from "./constructs/v1alpha2/relationship";

// v1alpha3
import {HttpsSchemasMeshplayRelationshipJson  as V1alpha3Relationship} from "./constructs/v1alpha3/relationship"
import { 
  HttpsSchemasMeshplaySelectorsJson as V1alpha3Selectors
} from "./constructs/v1alpha3/selectors"

// v1beta1
import { HttpsSchemasMeshplayComponentJson,Styles} from "./constructs/v1beta1/component";
import { HttpsSchemasMeshplayModelJson } from "./constructs/v1beta1/model";
import {DesignSchema as V1beta1DesignSchema} from "./constructs/v1beta1/designs"


//OpenAPI
import {components,paths,webhooks} from "./openapi"

import coreJson from "../schemas/constructs/core.json"

export namespace core {
  export type EntityStyles = Styles;
}

// Contructs 
export namespace v1alpha1 {
  export type CatalogData = CatalogDataT;
  export type CatalogCaveatsAndConsiderations = CaveatsAndConsiderations;
  export type Capability = HttpsSchemasMeshplayCapabilityJson;
}

export namespace v1alpha2 {
  export type Design = DesignSchema;
  export type Relationship = HttpsSchemasMeshplayRelationshipJson;
}

export namespace v1alpha3 {
  export type Relationship = V1alpha3Relationship;
  export type Selectors = V1alpha3Selectors;
}

export namespace v1beta1 {
  export type Component = HttpsSchemasMeshplayComponentJson;
  export type Model = HttpsSchemasMeshplayModelJson;
  export type Design = V1beta1DesignSchema;
}

// OpenAPI
export namespace api {
  export type Components = components;
  export type Paths = paths;
  export type Webhooks = webhooks;
}