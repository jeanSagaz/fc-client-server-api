IF NOT EXISTS (SELECT * FROM sys.databases WHERE name = 'fc')
BEGIN
  CREATE DATABASE fc;
END;
GO

USE fc
GO

IF OBJECT_ID('UsdBrl', 'U') IS NULL
BEGIN

  CREATE TABLE [dbo].[UsdBrl](
	[Id] [uniqueidentifier] NOT NULL,
	[Code] [nvarchar](10) NULL,
	[Codein] [nvarchar](10) NULL,
	[Name] [nvarchar](50) NULL,
	[High] [nvarchar](10) NULL,
	[Low] [nvarchar](10) NULL,
	[VarBid] [nvarchar](10) NULL,
	[PctChange] [nvarchar](10) NULL,
	[Bid] [nvarchar](10) NULL,
	[Ask] [nvarchar](10) NULL,
	[Timestamp] [nvarchar](15) NULL,
	[CreateDate] [nvarchar](25) NULL
 CONSTRAINT [PK_UsdBrl] PRIMARY KEY CLUSTERED 
(
	[Id] ASC
)WITH (PAD_INDEX = OFF, STATISTICS_NORECOMPUTE = OFF, IGNORE_DUP_KEY = OFF, ALLOW_ROW_LOCKS = ON, ALLOW_PAGE_LOCKS = ON) ON [PRIMARY]
) ON [PRIMARY]

END;
GO
