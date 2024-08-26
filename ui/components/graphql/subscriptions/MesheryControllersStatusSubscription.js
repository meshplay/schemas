import { graphql, requestSubscription } from 'react-relay';
import { createRelayEnvironment } from '../../../lib/relayEnvironment';

const meshplayControllersStatusSubscription = graphql`
  subscription MeshplayControllersStatusSubscription($k8scontextIDs: [String!]) {
    subscribeMeshplayControllersStatus(k8scontextIDs: $k8scontextIDs) {
      contextId
      controller
      status
    }
  }
`;

export default function subscribeMeshplayControllersStatus(dataCB, variables) {
  const environment = createRelayEnvironment({});
  return requestSubscription(environment, {
    subscription: meshplayControllersStatusSubscription,
    variables: { k8scontextIDs: variables },
    onNext: dataCB,
    onError: (error) => console.log(`An error occured:`, error),
  });
}
