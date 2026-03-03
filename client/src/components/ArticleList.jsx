import { useEffect, useState } from "react";
import { Box, Heading, Text, Stack } from "@chakra-ui/react";

export default function ArticleList({ refresh }) {
  const [articles, setArticles] = useState([]);

  useEffect(() => {
    fetch("http://localhost:8080/articles?limit=10&offset=0")
      .then((res) => res.json())
      .then((data) => setArticles(data));
  }, [refresh]); // re-fetch when refresh changes

  return (
    <Stack spacing={4} mt={6}>
      {articles.map((article) => (
        <Box key={article.ID} borderWidth="1px" borderRadius="md" p={4}>
          <Heading size="md">{article.Title}</Heading>
          <Text>{article.Content}</Text>
          <Text fontSize="sm" color="gray.500">
            Category: {article.Category} | Status: {article.Status}
          </Text>
        </Box>
      ))}
    </Stack>
  );
}
