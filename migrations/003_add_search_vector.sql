ALTER TABLE document_pages
ADD COLUMN search_vector tsvector;

UPDATE document_pages
SET search_vector = to_tsvector('spanish', content);

CREATE INDEX idx_document_pages_search
ON document_pages
USING GIN(search_vector);