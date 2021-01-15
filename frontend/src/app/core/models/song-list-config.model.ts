export interface SongListConfig {
    type: string;
  
    filters: {
        views?: string,
        release_date?: string
        limit?: number,
        offset?: number
    };
}
  