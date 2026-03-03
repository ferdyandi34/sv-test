import { useState } from "react";
import {
  Box,
  Input,
  Textarea,
  Select,
  Button,
  Heading,
  Stack,
} from "@chakra-ui/react";

export default function CreateArticleForm({ onCreated }) {
  const [title, setTitle] = useState("");
  const [content, setContent] = useState("");
  const [category, setCategory] = useState("");
  const [status, setStatus] = useState("Draft");

  const handleSubmit = async (e) => {
    e.preventDefault();
    const article = { title, content, category, status };
    await fetch("http://localhost:8080/articles", {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify(article),
    });
    setTitle("");
    setContent("");
    setCategory("");
    setStatus("Draft");
    onCreated(); // trigger refresh
  };

  return (
    <Box
      as="form"
      onSubmit={handleSubmit}
      p={4}
      borderWidth="1px"
      borderRadius="md"
    >
      <Heading size="md" mb={4}>
        Create New Article
      </Heading>
      <Stack spacing={3}>
        <Input
          placeholder="Title"
          value={title}
          onChange={(e) => setTitle(e.target.value)}
          required
        />
        <Textarea
          placeholder="Content"
          value={content}
          onChange={(e) => setContent(e.target.value)}
          required
        />
        <Input
          placeholder="Category"
          value={category}
          onChange={(e) => setCategory(e.target.value)}
        />
        <Select value={status} onChange={(e) => setStatus(e.target.value)}>
          <option value="Draft">Draft</option>
          <option value="Publish">Publish</option>
          <option value="Trash">Trash</option>
        </Select>
        <Button type="submit" colorScheme="blue">
          Create
        </Button>
      </Stack>
    </Box>
  );
}
