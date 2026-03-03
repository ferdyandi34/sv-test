import { ChakraProvider, Container, Heading } from "@chakra-ui/react";
import { useState } from "react";
import ArticleList from "./components/ArticleList";
import CreateArticleForm from "./components/CreateArticleForm";

export const BASE_URL = "http://localhost:5000/api";

function App() {
  const [refresh, setRefresh] = useState(0);

  return (
    <ChakraProvider>
      <Container maxW="container.md" py={8}>
        <Heading mb={6}>Articles</Heading>
        <CreateArticleForm onCreated={() => setRefresh(refresh + 1)} />
        <ArticleList refresh={refresh} />
      </Container>
    </ChakraProvider>
  );
}

export default App;

