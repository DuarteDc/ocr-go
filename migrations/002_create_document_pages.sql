CREATE TABLE document_pages (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),

    document_id UUID NOT NULL
        REFERENCES documents(id)
        ON DELETE CASCADE,

    page_number INTEGER NOT NULL,

    content TEXT NOT NULL,

    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),

    CONSTRAINT document_pages_unique_page
        UNIQUE (document_id, page_number)
);

CREATE INDEX idx_document_pages_document_id
    ON document_pages(document_id);
