-- campaigns table
CREATE TABLE IF NOT EXISTS campaigns (
    id TEXT PRIMARY KEY,
    name TEXT,
    img TEXT,
    cta TEXT,
    status TEXT CHECK (status IN ('ACTIVE', 'INACTIVE')) NOT NULL
);

-- targeting rules table
CREATE TABLE IF NOT EXISTS targeting_rules (
    id SERIAL PRIMARY KEY,
    campaign_id TEXT REFERENCES campaigns(id) ON DELETE CASCADE,
    include_app TEXT[],
    exclude_app TEXT[],
    include_os TEXT[],
    exclude_os TEXT[],
    include_country TEXT[],
    exclude_country TEXT[]
);

-- trigger to notify Redis listener (set up in a future step)
CREATE OR REPLACE FUNCTION notify_campaign_change()
RETURNS trigger AS $$
BEGIN
  PERFORM pg_notify('campaign_change', NEW.id);
  RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER campaign_update_trigger
AFTER INSERT OR UPDATE OR DELETE ON campaigns
FOR EACH ROW EXECUTE FUNCTION notify_campaign_change();
