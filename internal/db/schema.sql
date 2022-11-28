CREATE TABLE IF NOT EXISTS events ( 
        id serial primary key, 
        tag varchar (255) NOT NULL,
        organization varchar (255) NOT NULL,
        name varchar (255) NOT NULL, 
        available integer NOT NULL, 
        capacity integer NOT NULL, 
        status integer, 
        frontend text NOT NULL, 
        exercises text NOT NULL, 
        started_at timestamp NOT NULL,
        finish_expected timestamp NOT NULL, 
        finished_at timestamp NOT NULL, 
        createdBy text NOT NULL,
        secretKey text NOT NULL
);

CREATE TABLE IF NOT EXISTS teams (
        id serial primary key,
        tag varchar(255) NOT NULL,
        event_id integer NOT NULL,
        email varchar (255) NOT NULL,
        name varchar (255) NOT NULL, 
        password varchar (255) NOT NULL,
        created_at timestamp NOT NULL,
        last_access timestamp NOT NULL,
        solved_challenges text NOT NULL
);


-- Admin related tables
CREATE TABLE IF NOT EXISTS organizations (
        id serial primary key,
        name varchar (255) NOT NULL,
        owner_user varchar(255) NOT NULL,
        owner_email varchar(255) NOT NULL,
        UNIQUE(name)
);
CREATE UNIQUE INDEX orgname_lower_index ON organizations (LOWER(name));

CREATE TABLE IF NOT EXISTS profiles (
        id serial primary key, 
        name varchar (255) NOT NULL, 
        secret boolean NOT NULL, 
        organization varchar(255) NOT NULL REFERENCES organizations (name) ON DELETE CASCADE,
        challenges text NOT NULL
);        
CREATE UNIQUE INDEX profilename_lower_index ON profiles (LOWER(name));

CREATE TABLE IF NOT EXISTS admin_users (
        id serial primary key, 
        username varchar (255) NOT NULL, 
        password varchar (255) NOT NULL,
        full_name varchar (255) NOT NULL,
        email varchar (255) NOT NULL,
        role varchar (255) NOT NULL,
        organization varchar (255) NOT NULL REFERENCES organizations (name) ON DELETE CASCADE
);
CREATE UNIQUE INDEX username_lower_index ON Admin_users (LOWER(username));

-- TODO remove statelock from db
CREATE TABLE IF NOT EXISTS agents (
        id serial primary key,
        name varchar (255) NOT NULL,
        url varchar (255) NOT NULL,
        sign_key varchar (255) NOT NULL,
        auth_key varchar (255) NOT NULL,
        tls boolean NOT NULL DEFAULT true,
        statelock boolean NOT NULL DEFAULT false
);
CREATE UNIQUE INDEX agentname_lower_index ON agents (LOWER(name));

CREATE TABLE IF NOT EXISTS frontends (
        id serial primary key, 
        name varchar (255) NOT NULL,
        image varchar (255) NOT NULL,
        memoryMB integer
);
CREATE UNIQUE INDEX frontendname_lower_index ON frontends (LOWER(name));

